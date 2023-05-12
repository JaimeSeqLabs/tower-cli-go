package provider

import (
	"openapi"
	"os"

	"github.com/spf13/cobra"
)

func NewK8sCredentialsProvider(cmd *cobra.Command) CredentialsProvider {
	return k8sCredsProvider{cmd}
}

type k8sCredsProvider struct {
	Cmd *cobra.Command
}

func (p k8sCredsProvider) BaseUrl() string {
	return ""
}

func (p k8sCredsProvider) Type() CredentialsProviderEnum {
	return K8S
}

func (p k8sCredsProvider) SecurityKeys() (openapi.SecurityKeys, error) {

	certificatePath, err := p.Cmd.Flags().GetString("certificate")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	privateKeyPath, err := p.Cmd.Flags().GetString("private-key")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	token, err := p.Cmd.Flags().GetString("k8s-token")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	k := openapi.K8sSecurityKeys{}

	if certificatePath != "" && privateKeyPath != "" {
		certData, err := os.ReadFile(certificatePath)
		if err != nil {
			return openapi.SecurityKeys{}, err
		}
		pkData, err := os.ReadFile(privateKeyPath)
		if err != nil {
			return openapi.SecurityKeys{}, err
		}

		k.Certificate = string(certData)
		k.PrivateKey = string(pkData)
	} else {
		k.Token = token
	}

	return openapi.SecurityKeys{
		Discriminator: string(K8S),
		Certificate:   k.Certificate,
		PrivateKey:    k.PrivateKey,
		Token:         k.Token,
	}, nil
}
