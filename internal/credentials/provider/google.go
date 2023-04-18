package provider

import (
	"openapi"
	"os"

	"github.com/spf13/cobra"
)

func NewGoogleCredentialsProvider(cmd *cobra.Command) CredentialsProvider {
	return googleCredsProvider{cmd}
}

type googleCredsProvider struct {
	Cmd *cobra.Command
}

func (p googleCredsProvider) BaseUrl() string {
	return ""
}

func (p googleCredsProvider) Type() CredentialsProviderEnum {
	return GOOGLE
}

func (p googleCredsProvider) SecurityKeys() (openapi.SecurityKeys, error) {

	keyPath, err := p.Cmd.Flags().GetString("key")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	data, err := os.ReadFile(keyPath)
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	k := openapi.GoogleSecurityKeys{Data: string(data)}

	return openapi.SecurityKeys{
		Discriminator: string(GOOGLE),
		Data:          k.Data,
	}, nil
}
