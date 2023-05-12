package provider

import (
	"github.com/spf13/cobra"
	"openapi"
)

func NewGitLabCredentialsProvider(cmd *cobra.Command) CredentialsProvider {
	return gitLabCredsProvider{cmd}
}

type gitLabCredsProvider struct {
	Cmd *cobra.Command
}

func (p gitLabCredsProvider) BaseUrl() string {
	baseUrl, err := p.Cmd.Flags().GetString("base-url")
	if err != nil {
		return ""
	}
	return baseUrl
}

func (p gitLabCredsProvider) Type() CredentialsProviderEnum {
	return GITLAB
}

func (p gitLabCredsProvider) SecurityKeys() (openapi.SecurityKeys, error) {

	username, err := p.Cmd.Flags().GetString("username")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	password, err := p.Cmd.Flags().GetString("password")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	token, err := p.Cmd.Flags().GetString("gitlab-token")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	k := openapi.GitLabSecurityKeys{
		Username: username,
		Password: password,
		Token:    token,
	}

	return openapi.SecurityKeys{
		Discriminator: string(GITLAB),
		Username:      k.Username, // note 'n' lowercase
		Password:      k.Password,
		Token:         k.Token,
	}, nil
}
