package main

import (
	"github.com/spf13/cobra"

	"github.com/joshqu1985/lego/cmd/internal/client"
	"github.com/joshqu1985/lego/cmd/internal/config"
	"github.com/joshqu1985/lego/cmd/internal/server"
	"github.com/joshqu1985/lego/logs"
)

var rootCmd = &cobra.Command{
	Use: "lego",
}

func Init() {
	client.Init()
	rootCmd.AddCommand(client.Cmd)

	server.Init()
	rootCmd.AddCommand(server.Cmd)

	rootCmd.AddCommand(config.Cmd)
}

func main() {
	Init()

	if err := rootCmd.Execute(); err != nil {
		logs.Error("execute failed: %v", err)

		return
	}
}
