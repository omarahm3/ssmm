package api

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/omarahm3/ssmm/pkg/config"
)

func CreateSsmClient(config config.AwsConfig) *ssm.Client {
	var endpoint ssm.EndpointResolver

	if config.Endpoint == "" {
		endpoint = nil
	} else {
		endpoint = ssm.EndpointResolverFromURL(config.Endpoint)
	}

	return ssm.New(ssm.Options{
		Region:           *aws.String(config.Region),
		Credentials:      credentials.NewStaticCredentialsProvider(config.Access, config.Secret, ""),
		EndpointResolver: endpoint,
	})
}
