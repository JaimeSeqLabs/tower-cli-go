package provider

import (
	"openapi"

	"github.com/spf13/cobra"
)

func NewGitHubCredentialsProvider(cmd *cobra.Command) CredentialsProvider {
	return gitHubCredsProvider{cmd}
}

type gitHubCredsProvider struct {
	Cmd *cobra.Command
}

func (p gitHubCredsProvider) BaseUrl() string {
	return ""
}

func (p gitHubCredsProvider) Type() CredentialsProviderEnum {
	return GITHUB
}

func (p gitHubCredsProvider) SecurityKeys() (openapi.SecurityKeys, error) {

	username, err := p.Cmd.Flags().GetString("username")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	password, err := p.Cmd.Flags().GetString("password")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	k := openapi.GitHubSecurityKeys{
		Username: username,
		Password: password,
	}

	return openapi.SecurityKeys{
		Discriminator: string(GITHUB),
		Username:      k.Username, // note 'n' lowercase
		Password:      k.Password,
	}, nil
}
