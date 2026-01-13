package client

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/joshqu1985/lego/cmd/internal/pkg"
	"github.com/joshqu1985/lego/logs"
)

var (
	FlagClientType string
	FlagCodesDst   string
	FlagProtoFile  string

	Cmd = &cobra.Command{
		Use:   "client",
		Short: "Generate client code",
		Run:   run,
	}

	PackageName string
)

func Init() {
	Cmd.Flags().StringVarP(&FlagClientType, pkg.StringType, "t", "", "client type (http、grpc)")
	_ = Cmd.MarkFlagRequired(pkg.StringType)

	Cmd.Flags().StringVarP(&FlagCodesDst, pkg.StringDst, "d", "", "generated codes dir")
	_ = Cmd.MarkFlagRequired(pkg.StringDst)

	Cmd.Flags().StringVarP(&FlagProtoFile, pkg.StringProto, "", "", "proto file path")
	_ = Cmd.MarkFlagRequired(pkg.StringProto)
}

func run(_ *cobra.Command, args []string) {
	PackageName = filepath.Base(FlagCodesDst)
	FlagCodesDst = filepath.Dir(FlagCodesDst)

	if err := MkdirDirs(FlagCodesDst); err != nil {
		logs.Error(err)

		return
	}

	tree, err := pkg.Parse(FlagProtoFile)
	if err != nil {
		logs.Error(err)

		return
	}

	files := make([]pkg.File, 0)
	if FlagClientType == "http" {
		rfiles, xerr := GenRest(tree)
		if xerr != nil {
			return
		}
		files = append(files, rfiles...)
	} else {
		gfiles, xerr := GenGrpc(tree)
		if xerr != nil {
			return
		}
		files = append(files, gfiles...)
	}

	pkg.Writes(FlagCodesDst, files)
}

func GenRest(tree *pkg.Tree) ([]pkg.File, error) {
	files := make([]pkg.File, 0)
	if err := MkdirDirs(FlagCodesDst + "/" + PackageName); err != nil {
		logs.Error(err)

		return files, err
	}

	mfile, err := genRestModel(PackageName, FlagProtoFile, tree)
	if err != nil {
		logs.Error(err)

		return files, err
	}
	files = append(files, mfile)

	rfile, err := genRestRequest(PackageName, FlagProtoFile, tree)
	if err != nil {
		logs.Error(err)

		return files, err
	}

	files = append(files, rfile)

	return files, nil
}

func GenGrpc(tree *pkg.Tree) ([]pkg.File, error) {
	files := make([]pkg.File, 0)

	if err := genRpcProtoFile(PackageName, FlagProtoFile, FlagCodesDst); err != nil {
		logs.Error(err)

		return files, err
	}

	file, err := genRpcRequest(PackageName, FlagProtoFile, tree)
	if err != nil {
		logs.Error(err)

		return files, err
	}

	files = append(files, file)

	return files, nil
}

func MkdirDirs(dst string) error {
	if _, err := os.Stat(dst); err == nil {
		return nil
	}

	return pkg.Mkdir(dst)
}
