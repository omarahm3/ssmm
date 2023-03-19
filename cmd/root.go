package cmd

import (
	"fmt"
	"os"

	"github.com/omarahm3/ssmm/pkg/config"
	"github.com/omarahm3/ssmm/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	configFile         string
	ssmConfig          *config.Config
	needConfigCommands = []string{"apply", "list"}
	rootCmd            = &cobra.Command{
		Use:   "ssmm",
		Short: "manage multiple projects environment variables in SSM Parameter Store",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if !isNeedConfigCommand(cmd.Name()) {
				return
			}

			if configFile == "" {
				utils.FatalPrint("config file is required")
			}

			c, err := config.Load(configFile)
			utils.CheckError(err)
			ssmConfig = c
		},
	}
)

func isNeedConfigCommand(cmd string) bool {
	for _, c := range needConfigCommands {
		if c == cmd {
			return true
		}
	}
	return false
}

func Init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
