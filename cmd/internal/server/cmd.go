package server

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang/glog"
	"github.com/jinzhu/inflection"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joshqu1985/lego/cmd/internal/pkg"
)

const (
	SQL_FOMATE_TABLES = "SELECT TABLE_NAME as name, TABLE_COMMENT as comment FROM" +
		" information_schema.TABLES WHERE table_schema='%s'"
	SQL_FOMATE_TABLE_DETAIL = "SELECT COLUMN_NAME as name, DATA_TYPE as data_type, " +
		" COLUMN_COMMENT as comment FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA='%s' AND table_name='%s'"

	SERVER_TYPE_HTTP = "http"
	SERVER_TYPE_GRPC = "grpc"
)

var (
	FlagServerType string
	FlagCodesDst   string

	FlagProtoFile     string
	FlagMysqlHost     string
	FlagMysqlDatabase string
	FlagMysqlAuth     string

	PackageName string

	Cmd = &cobra.Command{
		Use:   "server",
		Short: "Generate server code",
		Run:   run,
	}
)

func Init() {
	Cmd.Flags().StringVarP(&FlagServerType, pkg.StringType, "t", "", "codes type (http, grpc, dbstructs)")
	_ = Cmd.MarkFlagRequired(pkg.StringType)

	Cmd.Flags().StringVarP(&FlagCodesDst, pkg.StringDst, "d", "", "generated codes dir")
	_ = Cmd.MarkFlagRequired(pkg.StringDst)

	Cmd.Flags().StringVarP(&FlagProtoFile, pkg.StringProto, "", "", "proto file path")
	Cmd.Flags().StringVarP(&FlagMysqlHost, pkg.StringHost, "", "", "mysql enpoint addr:port")
	Cmd.Flags().StringVarP(&FlagMysqlDatabase, pkg.StringDB, "", "", "mysql database")
	Cmd.Flags().StringVarP(&FlagMysqlAuth, pkg.StringAuth, "", "", "mysql username:password")
}

func run(_ *cobra.Command, args []string) {
	if err := MkdirDirs(FlagCodesDst); err != nil {
		glog.Error(err)

		return
	}

	PackageName = filepath.Base(FlagCodesDst)
	FlagCodesDst = filepath.Dir(FlagCodesDst)

	tree, err := pkg.Parse(FlagProtoFile)
	if err != nil {
		glog.Error(err)

		return
	}

	codeFiles := make([]pkg.File, 0)
	switch FlagServerType {
	case SERVER_TYPE_HTTP:
		file, xerr := GenRestMain(PackageName)
		if xerr != nil {
			glog.Fatal(xerr)

			return
		}
		codeFiles = append(codeFiles, file)

		files, yerr := GenRestApis(PackageName, FlagProtoFile, tree)
		if yerr != nil {
			glog.Fatal(yerr)

			return
		}
		codeFiles = append(codeFiles, files...)
	case SERVER_TYPE_GRPC:
		file, xerr := GenRpcMain(PackageName)
		if xerr != nil {
			glog.Fatal(xerr)

			return
		}
		codeFiles = append(codeFiles, file)

		files, yerr := GenRpcApis(PackageName, FlagProtoFile, tree)
		if yerr != nil {
			glog.Fatal(yerr)

			return
		}
		codeFiles = append(codeFiles, files...)

		files, zerr := GenRpcProtos(PackageName, FlagProtoFile, FlagCodesDst, tree)
		if zerr != nil {
			glog.Fatal(zerr)

			return
		}
		codeFiles = append(codeFiles, files...)
	case "dbstructs":
		tables, xerr := ReadMysqlTables(FlagMysqlHost, FlagMysqlDatabase, FlagMysqlAuth)
		if xerr != nil {
			glog.Fatal(xerr)

			return
		}
		files, yerr := GenModelEntities(PackageName, tables)
		if yerr != nil {
			glog.Fatal(yerr)

			return
		}
		codeFiles = append(codeFiles, files...)
	}

	if FlagServerType == SERVER_TYPE_HTTP || FlagServerType == SERVER_TYPE_GRPC {
		{
			file, xerr := GenConfig(PackageName)
			if xerr != nil {
				glog.Fatal(xerr)

				return
			}
			codeFiles = append(codeFiles, file)
		}

		{
			files, yerr := GenModels(PackageName, FlagProtoFile, tree)
			if yerr != nil {
				glog.Fatal(yerr)

				return
			}
			codeFiles = append(codeFiles, files...)
		}

		{
			files, zerr := GenServices(PackageName, FlagProtoFile, tree)
			if zerr != nil {
				glog.Fatal(zerr)

				return
			}
			codeFiles = append(codeFiles, files...)
		}

		{
			file, kerr := GenRepository(PackageName)
			if kerr != nil {
				glog.Fatal(kerr)

				return
			}
			codeFiles = append(codeFiles, file)
		}
	}
	pkg.Writes(FlagCodesDst, codeFiles)
}

func MkdirDirs(dst string) error {
	if _, err := os.Stat(dst); err != nil {
		if xerr := pkg.Mkdir(dst); xerr != nil {
			return xerr
		}
	}
	if _, err := os.Stat(dst + "/cmd"); err != nil {
		if xerr := pkg.Mkdir(dst + "/cmd"); xerr != nil {
			return xerr
		}
	}
	if _, err := os.Stat(dst + "/internal/config"); err != nil {
		if xerr := pkg.Mkdir(dst + "/internal/config"); xerr != nil {
			return xerr
		}
	}
	if _, err := os.Stat(dst + "/internal/model/entity"); err != nil {
		if xerr := pkg.Mkdir(dst + "/internal/model/entity"); xerr != nil {
			return xerr
		}
	}
	if _, err := os.Stat(dst + "/internal/service"); err != nil {
		if xerr := pkg.Mkdir(dst + "/internal/service"); xerr != nil {
			return xerr
		}
	}
	if _, err := os.Stat(dst + "/internal/repo"); err != nil {
		if xerr := pkg.Mkdir(dst + "/internal/repo"); xerr != nil {
			return xerr
		}
	}
	if FlagServerType == SERVER_TYPE_HTTP {
		if err := pkg.Mkdir(dst + "/httpserver"); err != nil {
			return err
		}
	} else {
		if err := pkg.Mkdir(dst + "/rpcserver/protocols"); err != nil {
			return err
		}
	}

	return nil
}

func ReadMysqlTables(host, database, auth string) ([]*Table, error) {
	dsn := fmt.Sprintf("%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		auth, host, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 查询数据库所有表信息（表名 - 表描述）
	items := make([]*Table, 0)
	if xerr := db.Raw(fmt.Sprintf(SQL_FOMATE_TABLES, database)).Find(&items).Error; xerr != nil {
		return nil, xerr
	}

	for _, item := range items {
		if item.Name == "gorp_migrations" {
			continue
		}
		item.SingularName = inflection.Singular(item.Name)
		item.PluralName = inflection.Plural(item.SingularName)
	}

	// 查询每一个表的详细信息
	for _, item := range items {
		if xerr := db.Raw(fmt.Sprintf(SQL_FOMATE_TABLE_DETAIL, database, item.Name)).
			Find(&item.Columns).Error; xerr != nil {
			return nil, xerr
		}
	}

	return items, nil
}
