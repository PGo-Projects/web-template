package main

import (
	"github.com/PGo-Projects/output"
	"github.com/PGo-Projects/web-template/internal/config"
	"github.com/PGo-Projects/web-template/internal/server"
	"github.com/spf13/cobra"
)

var (
	// TODO: Insert cobra command name here
	ServerCmd = &cobra.Command{
		Use: "<INSERT NAME HERE>",
		Run: server.Run,
	}
)

func init() {
	ServerCmd.PersistentFlags().StringVar(&config.Filename, "config", "",
		"config file (default is config.toml)")
	ServerCmd.PersistentFlags().BoolVar(&config.DevRun, "dev", false,
		"Run the server on a dev machine")
	cobra.OnInitialize(config.Init)
}

func main() {
	if err := ServerCmd.Execute(); err != nil {
		output.ErrorAndExit(err, 1)
	}
}
