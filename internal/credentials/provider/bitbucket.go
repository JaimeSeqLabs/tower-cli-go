package provider

import (
	"github.com/spf13/cobra"
	"openapi"
)

func NewBitBucketCredentialsProvider(cmd *cobra.Command) CredentialsProvider {
	return bitBucketCredsProvider{cmd}
}

type bitBucketCredsProvider struct {
	Cmd *cobra.Command
}

func (p bitBucketCredsProvider) BaseUrl() string {
	return ""
}

func (p bitBucketCredsProvider) Type() CredentialsProviderEnum {
	return BITBUCKET
}

func (p bitBucketCredsProvider) SecurityKeys() (openapi.SecurityKeys, error) {

	username, err := p.Cmd.Flags().GetString("username")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	password, err := p.Cmd.Flags().GetString("password")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	k := openapi.BitBucketSecurityKeys{
		Username: username,
		Password: password,
	}

	return openapi.SecurityKeys{
		Discriminator: string(BITBUCKET),
		Username:      k.Username, // note 'n' lowercase
		Password:      k.Password,
	}, nil
}
