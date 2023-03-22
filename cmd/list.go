package cmd

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/omarahm3/ssmm/pkg/api"
	"github.com/omarahm3/ssmm/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	listCommand = &cobra.Command{
		Use:   "list",
		Short: "list current parameters in SSM Parameter Store",
		Run:   runListCommand,
	}
)

func runListCommand(cmd *cobra.Command, args []string) {
	client := api.CreateSsmClient(ssmConfig.Aws)

	parameters, err := api.GetAllParameters(client)
	utils.CheckError(err)

	for i, parameter := range parameters {
		printParameter(i, parameter)
	}
}

func printParameter(index int, parameter types.ParameterMetadata) {
	s := strings.Split(strings.TrimSpace(*parameter.Name), "/")

	if len(s) != 4 {
		return
	}

	fmt.Printf("%d) [ %s ] - [ %s ]\n", index+1, *parameter.Name, *&parameter.Type)
}

func init() {
	rootCmd.AddCommand(listCommand)
}
