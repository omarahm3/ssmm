# ssmm - AWS SSM Parameter Store Manager

Simple tool to manage SSM Parameter Store parameters for multiple projects with the goal of having a single source of truth for all parameters per project or to combine them all in single config file.

Parameters are stored with respect to there project and environment, so having:

```yaml
projects:
  - name: "sqrl"
    environments:
      - name: "local"
        variables:
          - key: "NODE_ENV"
            value: "development"
            overwrite: true
            securet: true
```

will result in the following parameter:

```json
{
    "Name": "/sqrl/local/NODE_ENV",
    "Type": "String",
    "LastModifiedDate": 1679338581.762,
    "LastModifiedUser": "N/A",
    "Version": 6,
    "DataType": "text"
}
```

## Usage

```bash
❯ ssmm help

manage multiple projects environment variables in SSM Parameter Store

Usage:
  ssmm [command]

Available Commands:
  apply       apply all environment variables to SSM Parameter Store (SSM)
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        list current parameters in SSM Parameter Store

Flags:
  -c, --config string   config file
  -h, --help            help for ssmm
```


## Config file

The simplest config file should go as follow:

```yaml
aws: # optional
  access: "test" #  optional - aws access key or AWS_ACCESS_KEY_ID
  secret: "test" # optional - aws secret key or AWS_SECRET_ACCESS_KEY
  region: "us-east-1" # optional - aws region or AWS_DEFAULT_REGION
  endpoint: "http://localhost:4566" # optional - use it if you want to test with localstack or AWS_ENDPOINT

projects:
  - name: "sqrl"
    environments:
      - name: "local"
        variables:
          - key: "NODE_ENV"
            value: "development"
            overwrite: true # optional - this will overwrite the value if it already exists
          - key: "PORT"
            value: 3000
            overwrite: true
      - name: "production"
        variables:
          - key: "NODE_ENV"
            value: "production"
          - key: "PORT"
            value: 80
            overwrite: true
          - key: "DATABASE_URL"
            value: "postgres://user:pass@localhost:5432/db"
            secure: true
```

## Run


```bash
❯ ssmm apply -f example.yaml
❯ AWS_ACCESS_KEY_ID=test1 ssmm apply -f example.yaml
❯ AWS_ACCESS_KEY_ID=test1 AWS_SECRET_ACCESS_KEY=test2 AWS_DEFAULT_REGION=us-east-2 AWS_ENDPOINT=http://localhost:4567 ssmm apply -f example.yaml
```

**Note:** Passing AWS variables to the CLI will override aws config block in the config file.

## Install

If you already have go installed then:

```bash
go install github.com/omarahm3/ssm@latest
```

Or from releases page
