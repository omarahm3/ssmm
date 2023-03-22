package api

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/omarahm3/ssmm/pkg/config"
)

func Play(client *ssm.Client) {
	out, err := client.DescribeParameters(context.Background(), &ssm.DescribeParametersInput{})
	if err != nil {
		panic(err)
	}

	fmt.Println(out.Parameters)
}

func GetAllParameters(client *ssm.Client) ([]types.ParameterMetadata, error) {
	var parameters []types.ParameterMetadata

	out, err := client.DescribeParameters(context.Background(), &ssm.DescribeParametersInput{})
	if err != nil {
		return nil, err
	}

	parameters = append(parameters, out.Parameters...)

	for out.NextToken != nil {
		out, err = client.DescribeParameters(context.Background(), &ssm.DescribeParametersInput{
			NextToken: out.NextToken,
		})
		if err != nil {
			return nil, err
		}
		parameters = append(parameters, out.Parameters...)
	}

	return parameters, nil
}

func AddProjectVariables(client *ssm.Client, project config.Project) error {
	var parameters []*ssm.PutParameterInput

	for _, env := range project.Environments {
		for _, variable := range env.Variables {
			variablePath := fmt.Sprintf("/%s/%s/%s", project.Name, env.Name, variable.Key)
			variableType := types.ParameterTypeString

			if variable.Secure {
				variableType = types.ParameterTypeSecureString
			}

			parameters = append(parameters, &ssm.PutParameterInput{
				Name:      aws.String(variablePath),
				Value:     aws.String(variable.Value),
				Type:      variableType,
				Overwrite: aws.Bool(variable.Overwrite),
			})
		}
	}

	for _, param := range parameters {
		fmt.Printf("> adding new variable [ %s ] with value [ %s ]\n", *param.Name, *param.Value)
		_, err := client.PutParameter(context.Background(), param)
		if err != nil && strings.Contains(err.Error(), "ParameterAlreadyExists") {
			fmt.Printf("X> parameter [ %s ] already exists, ignoring...\n", *param.Name)
			continue
		}

		if err != nil {
			return err
		}
	}

	return nil
}
