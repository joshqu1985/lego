package config

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "config",
	Short: "Generate config file",
	Run:   run,
}

func run(_ *cobra.Command, args []string) {
}
