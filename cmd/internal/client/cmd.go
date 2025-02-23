package client

import (
	"os"
	"path/filepath"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/joshqu1985/lego/cmd/internal/pkg"
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

func init() {
	Cmd.Flags().StringVarP(&FlagClientType, "type", "t", "", "client type (http、grpc)")
	Cmd.MarkFlagRequired("type")

	Cmd.Flags().StringVarP(&FlagCodesDst, "dst", "d", "", "generated codes dir")
	Cmd.MarkFlagRequired("dst")

	Cmd.Flags().StringVarP(&FlagProtoFile, "proto", "", "", "proto file path")
	Cmd.MarkFlagRequired("proto")
}

func run(_ *cobra.Command, args []string) {
	PackageName = filepath.Base(FlagCodesDst)
	FlagCodesDst = filepath.Dir(FlagCodesDst)

	if err := MkdirDirs(FlagCodesDst); err != nil {
		glog.Error(err)
		return
	}

	tree, err := pkg.Parse(FlagProtoFile)
	if err != nil {
		glog.Error(err)
		return
	}

	files := []pkg.File{}
	if FlagClientType == "http" {
		if err := MkdirDirs(FlagCodesDst + "/" + PackageName); err != nil {
			glog.Error(err)
			return
		}

		mfile, err := genRestModel(PackageName, FlagProtoFile, tree)
		if err != nil {
			glog.Error(err)
			return
		}
		files = append(files, mfile)

		rfile, err := genRestRequest(PackageName, FlagProtoFile, tree)
		if err != nil {
			glog.Error(err)
			return
		}
		files = append(files, rfile)
	} else {
		if err := genRpcProtoFile(PackageName, FlagProtoFile, FlagCodesDst); err != nil {
			glog.Error(err)
			return
		}

		file, err := genRpcRequest(PackageName, FlagProtoFile, tree)
		if err != nil {
			glog.Error(err)
			return
		}
		files = append(files, file)
	}
	// TODO _test.go .md

	pkg.Writes(FlagCodesDst, files)
}

func MkdirDirs(dst string) error {
	if _, err := os.Stat(dst); err == nil {
		return nil
	}
	return pkg.Mkdir(dst)
}
