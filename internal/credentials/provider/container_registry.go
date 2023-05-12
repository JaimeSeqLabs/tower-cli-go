package provider

import (
	"openapi"

	"github.com/spf13/cobra"
)

func NewContainerRegistryCredentialsProvider(cmd *cobra.Command) CredentialsProvider {
	return azureCredsProvider{cmd}
}

type cRegCredsProvider struct {
	Cmd *cobra.Command
}

func (p cRegCredsProvider) BaseUrl() string {
	return ""
}

func (p cRegCredsProvider) Type() CredentialsProviderEnum {
	return CONTAINER_REG
}

func (p cRegCredsProvider) SecurityKeys() (openapi.SecurityKeys, error) {

	username, err := p.Cmd.Flags().GetString("username")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	password, err := p.Cmd.Flags().GetString("password")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	registry, err := p.Cmd.Flags().GetString("registry")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	k := openapi.ContainerRegistryKeys{
		UserName: username,
		Password: password,
		Registry: registry,
	}

	return openapi.SecurityKeys{
		Discriminator: string(CONTAINER_REG),
		UserName:      k.UserName, // note 'N' uppercase
		Password:      k.Password,
		Registry:      k.Registry,
	}, nil
}
