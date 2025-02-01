package server

import (
	"os"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/joshqu1985/lego/cmd/internal/pkg"
)

var (
	FlagProtoSrc   string
	FlagCodesDst   string
	FlagServerType string
	FlagServerName string

	Cmd = &cobra.Command{
		Use:   "server",
		Short: "Generate server code",
		Run:   run,
	}
)

func init() {
	Cmd.Flags().StringVarP(&FlagProtoSrc, "src", "s", "", "proto file path")
	Cmd.MarkFlagRequired("src")

	Cmd.Flags().StringVarP(&FlagCodesDst, "dst", "d", "", "generated codes dir")
	Cmd.MarkFlagRequired("dst")

	Cmd.Flags().StringVarP(&FlagServerType, "type", "t", "", "server type (rest、rpc)")
	Cmd.MarkFlagRequired("type")

	Cmd.Flags().StringVarP(&FlagServerName, "name", "n", "", "server name (folder name、go.mod name、registry name)")
	Cmd.MarkFlagRequired("name")
}

func run(_ *cobra.Command, args []string) {
	mkdirDirs()

	tree, err := pkg.Parse(FlagProtoSrc)
	if err != nil {
		glog.Fatal(err)
	}

	codeFiles := []pkg.File{}
	if FlagServerType == "rest" {
		files, err := GenRestApis(FlagServerName, FlagProtoSrc, tree)
		if err != nil {
			glog.Fatal(err)
			return
		}
		codeFiles = append(codeFiles, files...)
	} else {
		files, err := GenRpcApis(FlagServerName, FlagProtoSrc, tree)
		if err != nil {
			glog.Fatal(err)
			return
		}
		codeFiles = append(codeFiles, files...)
	}

	{
		files, err := GenModels(FlagServerName, FlagProtoSrc, tree)
		if err != nil {
			glog.Fatal(err)
			return
		}
		codeFiles = append(codeFiles, files...)
	}

	{
		files, err := GenServices(FlagServerName, FlagProtoSrc, tree)
		if err != nil {
			glog.Fatal(err)
			return
		}
		codeFiles = append(codeFiles, files...)
	}

	pkg.Writes(FlagCodesDst, codeFiles)
}

func mkdirDirs() {
	if _, err := os.Stat(FlagCodesDst + "/" + FlagServerName); err != nil {
		if err := pkg.Mkdir(FlagCodesDst + "/" + FlagServerName); err != nil {
			glog.Fatal(err)
		}
	}
	if _, err := os.Stat(FlagCodesDst + "/" + FlagServerName + "/internal/model"); err != nil {
		if err := pkg.Mkdir(FlagCodesDst + "/" + FlagServerName + "/internal/model"); err != nil {
			glog.Fatal(err)
		}
	}
	if _, err := os.Stat(FlagCodesDst + "/" + FlagServerName + "/internal/model/entity"); err != nil {
		if err := pkg.Mkdir(FlagCodesDst + "/" + FlagServerName + "/internal/model/entity"); err != nil {
			glog.Fatal(err)
		}
	}
	if _, err := os.Stat(FlagCodesDst + "/" + FlagServerName + "/internal/service"); err != nil {
		if err := pkg.Mkdir(FlagCodesDst + "/" + FlagServerName + "/internal/service"); err != nil {
			glog.Fatal(err)
		}
	}
	if _, err := os.Stat(FlagCodesDst + "/" + FlagServerName + "/internal/repo"); err != nil {
		if err := pkg.Mkdir(FlagCodesDst + "/" + FlagServerName + "/internal/repo"); err != nil {
			glog.Fatal(err)
		}
	}
	if FlagServerType == "rest" {
		if err := pkg.Mkdir(FlagCodesDst + "/" + FlagServerName + "/httpserver"); err != nil {
			glog.Fatal(err)
		}
	} else {
		if err := pkg.Mkdir(FlagCodesDst + "/" + FlagServerName + "/rpcserver"); err != nil {
			glog.Fatal(err)
		}
	}
}
