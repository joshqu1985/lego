package main

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/joshqu1985/lego/cmd/internal/client"
	"github.com/joshqu1985/lego/cmd/internal/config"
	"github.com/joshqu1985/lego/cmd/internal/server"
)

var rootCmd = &cobra.Command{
	Use: "lego",
}

func init() {
	rootCmd.AddCommand(client.Cmd)
	rootCmd.AddCommand(server.Cmd)
	rootCmd.AddCommand(config.Cmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		glog.Fatal(err)
	}
}
