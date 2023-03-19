# ssmm - AWS SSM Parameter Store Manager

A simple tool to manage all AWS SSM Parameter Store pparameters, with the goal of having a single source of truth for all parameters.

```
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
aws:
  access: "test" # aws access key
  secret: "test" # aws secret key
  region: "us-east-1"
  endpoint: "http://localhost:4566" # optional - use it if you want to test with localstack

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

## Install

If you already have go installed then:

```bash
go install github.com/omarahm3/ssm@latest
```

Or from releases page
