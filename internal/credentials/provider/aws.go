package provider

import (
	"github.com/spf13/cobra"
	"openapi"
)

func NewAwsCredentialsProvider(cmd *cobra.Command) CredentialsProvider {
	return awsCredsProvider{cmd}
}

type awsCredsProvider struct {
	Cmd *cobra.Command
}

func (p awsCredsProvider) BaseUrl() string {
	return ""
}

func (p awsCredsProvider) Type() CredentialsProviderEnum {
	return AWS
}

func (p awsCredsProvider) SecurityKeys() (openapi.SecurityKeys, error) {

	accessKey, err := p.Cmd.Flags().GetString("access-key")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	secretKey, err := p.Cmd.Flags().GetString("secret-key")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	arnRole, err := p.Cmd.Flags().GetString("assume-role-arn")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	k := openapi.AwsSecurityKeys{
		AccessKey:     accessKey,
		SecretKey:     secretKey,
		AssumeRoleArn: arnRole,
	}

	return openapi.SecurityKeys{
		Discriminator: string(AWS),
		AccessKey:     k.AccessKey,
		SecretKey:     k.SecretKey,
		AssumeRoleArn: k.AssumeRoleArn,
	}, nil
}
