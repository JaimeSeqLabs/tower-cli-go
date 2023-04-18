package provider

import (
	"openapi"
	"os"

	"github.com/spf13/cobra"
)

func NewSSHCredentialsProvider(cmd *cobra.Command) CredentialsProvider {
	return sshCredsProvider{cmd}
}

type sshCredsProvider struct {
	Cmd *cobra.Command
}

func (p sshCredsProvider) BaseUrl() string {
	return ""
}

func (p sshCredsProvider) Type() CredentialsProviderEnum {
	return SSH
}

func (p sshCredsProvider) SecurityKeys() (openapi.SecurityKeys, error) {

	keyPath, err := p.Cmd.Flags().GetString("key")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	passPhrase, _ := p.Cmd.Flags().GetString("passphrase")

	data, err := os.ReadFile(keyPath)
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	k := openapi.SshSecurityKeys{
		PrivateKey: string(data),
		Passphrase: passPhrase,
	}

	return openapi.SecurityKeys{
		Discriminator: string(SSH),
		PrivateKey:    k.PrivateKey,
		Passphrase:    k.Passphrase,
	}, nil
}
