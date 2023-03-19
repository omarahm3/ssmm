package cmd

import (
	"github.com/omarahm3/ssmm/pkg/api"
	"github.com/omarahm3/ssmm/pkg/config"
	"github.com/omarahm3/ssmm/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	applyCommand = &cobra.Command{
		Use:   "apply",
		Short: "apply all environment variables to SSM Parameter Store (SSM)",
		Run:   runApplyCommand,
	}
)

func runApplyCommand(cmd *cobra.Command, args []string) {
	c, err := config.Load(configFile)
	utils.CheckError(err)

	client := api.CreateSsmClient(c.Aws)

	for _, project := range c.Projects {
		api.AddProjectVariables(client, project)
	}
}

func init() {
	rootCmd.AddCommand(applyCommand)
}
