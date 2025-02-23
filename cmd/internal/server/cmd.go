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

func init() {
	Cmd.Flags().StringVarP(&FlagServerType, "type", "t", "", "codes type (http, grpc, dbstructs)")
	Cmd.MarkFlagRequired("type")

	Cmd.Flags().StringVarP(&FlagCodesDst, "dst", "d", "", "generated codes dir")
	Cmd.MarkFlagRequired("dst")

	Cmd.Flags().StringVarP(&FlagProtoFile, "proto", "", "", "proto file path")
	Cmd.Flags().StringVarP(&FlagMysqlHost, "host", "", "", "mysql enpoint addr:port")
	Cmd.Flags().StringVarP(&FlagMysqlDatabase, "db", "", "", "mysql database")
	Cmd.Flags().StringVarP(&FlagMysqlAuth, "auth", "", "", "mysql username:password")
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

	codeFiles := []pkg.File{}
	if FlagServerType == "http" {
		file, err := GenRestMain(PackageName)
		if err != nil {
			glog.Fatal(err)
			return
		}
		codeFiles = append(codeFiles, file)

		files, err := GenRestApis(PackageName, FlagProtoFile, tree)
		if err != nil {
			glog.Fatal(err)
			return
		}
		codeFiles = append(codeFiles, files...)
	} else if FlagServerType == "grpc" {
		file, err := GenRpcMain(PackageName)
		if err != nil {
			glog.Fatal(err)
			return
		}
		codeFiles = append(codeFiles, file)

		files, err := GenRpcApis(PackageName, FlagProtoFile, tree)
		if err != nil {
			glog.Fatal(err)
			return
		}
		codeFiles = append(codeFiles, files...)

		files, err = GenRpcProtos(PackageName, FlagProtoFile, FlagCodesDst, tree)
		if err != nil {
			glog.Fatal(err)
			return
		}
		codeFiles = append(codeFiles, files...)
	} else if FlagServerType == "dbstructs" {
		tables, err := ReadMysqlTables(FlagMysqlHost, FlagMysqlDatabase, FlagMysqlAuth)
		if err != nil {
			glog.Fatal(err)
			return
		}
		files, err := GenModelEntities(PackageName, tables)
		if err != nil {
			glog.Fatal(err)
			return
		}
		codeFiles = append(codeFiles, files...)
	}

	if FlagServerType == "http" || FlagServerType == "grpc" {
		{
			file, err := GenConfig(PackageName)
			if err != nil {
				glog.Fatal(err)
				return
			}
			codeFiles = append(codeFiles, file)
		}

		{
			files, err := GenModels(PackageName, FlagProtoFile, tree)
			if err != nil {
				glog.Fatal(err)
				return
			}
			codeFiles = append(codeFiles, files...)
		}

		{
			files, err := GenServices(PackageName, FlagProtoFile, tree)
			if err != nil {
				glog.Fatal(err)
				return
			}
			codeFiles = append(codeFiles, files...)
		}

		{
			file, err := GenRepository(PackageName)
			if err != nil {
				glog.Fatal(err)
				return
			}
			codeFiles = append(codeFiles, file)
		}
	}
	pkg.Writes(FlagCodesDst, codeFiles)
}

func MkdirDirs(dst string) error {
	if _, err := os.Stat(dst); err != nil {
		if err := pkg.Mkdir(dst); err != nil {
			return err
		}
	}
	if _, err := os.Stat(dst + "/cmd"); err != nil {
		if err := pkg.Mkdir(dst + "/cmd"); err != nil {
			return err
		}
	}
	if _, err := os.Stat(dst + "/internal/config"); err != nil {
		if err := pkg.Mkdir(dst + "/internal/config"); err != nil {
			return err
		}
	}
	if _, err := os.Stat(dst + "/internal/model/entity"); err != nil {
		if err := pkg.Mkdir(dst + "/internal/model/entity"); err != nil {
			return err
		}
	}
	if _, err := os.Stat(dst + "/internal/service"); err != nil {
		if err := pkg.Mkdir(dst + "/internal/service"); err != nil {
			return err
		}
	}
	if _, err := os.Stat(dst + "/internal/repo"); err != nil {
		if err := pkg.Mkdir(dst + "/internal/repo"); err != nil {
			return err
		}
	}
	if FlagServerType == "http" {
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

const (
	SQL_FOMATE_TABLES       = "SELECT TABLE_NAME as name, TABLE_COMMENT as comment FROM information_schema.TABLES WHERE table_schema='%s'"
	SQL_FOMATE_TABLE_DETAIL = "SELECT COLUMN_NAME as name, DATA_TYPE as data_type, COLUMN_COMMENT as comment FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA='%s' AND table_name='%s'"
)

func ReadMysqlTables(host, database, auth string) ([]*Table, error) {
	dsn := fmt.Sprintf("%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		auth, host, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 查询数据库所有表信息（表名 - 表描述）
	items := []*Table{}
	if err := db.Raw(fmt.Sprintf(SQL_FOMATE_TABLES, database)).Find(&items).Error; err != nil {
		return nil, err
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
		if err := db.Raw(fmt.Sprintf(SQL_FOMATE_TABLE_DETAIL, database, item.Name)).
			Find(&item.Columns).Error; err != nil {
			return nil, err
		}
	}
	return items, nil
}
