package client

import (
	"os"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/joshqu1985/lego/cmd/internal/pkg"
)

var (
	FlagProtoSrc   string
	FlagCodesDst   string
	FlagClientType string
	FlagPackage    string

	Cmd = &cobra.Command{
		Use:   "client",
		Short: "Generate client code",
		Run:   run,
	}
)

func init() {
	Cmd.Flags().StringVarP(&FlagProtoSrc, "src", "s", "", "proto file path")
	Cmd.MarkFlagRequired("src")

	Cmd.Flags().StringVarP(&FlagCodesDst, "dst", "d", "", "generated codes dir")
	Cmd.MarkFlagRequired("dst")

	Cmd.Flags().StringVarP(&FlagClientType, "type", "t", "", "client type (rest、rpc)")
	Cmd.MarkFlagRequired("type")

	Cmd.Flags().StringVarP(&FlagPackage, "package", "p", "", "client package name")
	Cmd.MarkFlagRequired("package")
}

func run(_ *cobra.Command, args []string) {
	if _, err := os.Stat(FlagCodesDst + "/" + FlagPackage); err != nil {
		if err := pkg.Mkdir(FlagCodesDst + "/" + FlagPackage); err != nil {
			glog.Fatal(err)
		}
	}

	tree, err := pkg.Parse(FlagProtoSrc)
	if err != nil {
		glog.Fatal(err)
	}

	files := []pkg.File{}
	if FlagClientType == "rest" {
		mfile, err := genRestModel(FlagPackage, FlagProtoSrc, tree)
		if err != nil {
			glog.Fatal(err)
			return
		}
		files = append(files, mfile)

		rfile, err := genRestRequest(FlagPackage, FlagProtoSrc, tree)
		if err != nil {
			glog.Fatal(err)
			return
		}
		files = append(files, rfile)
	} else {
		if err := genRpcProtoFile(FlagPackage, FlagProtoSrc, FlagCodesDst); err != nil {
			glog.Fatal(err)
			return
		}

		file, err := genRpcRequest(FlagPackage, FlagProtoSrc, tree)
		if err != nil {
			glog.Fatal(err)
			return
		}
		files = append(files, file)
	}
	// TODO _test.go .md

	pkg.Writes(FlagCodesDst, files)
}
