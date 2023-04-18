package provider

import (
	"openapi"

	"github.com/spf13/cobra"
)

func NewAzureCredentialsProvider(cmd *cobra.Command) CredentialsProvider {
	return azureCredsProvider{cmd}
}

type azureCredsProvider struct {
	Cmd *cobra.Command
}

func (p azureCredsProvider) BaseUrl() string {
	return ""
}

func (p azureCredsProvider) Type() CredentialsProviderEnum {
	return AZURE
}

func (p azureCredsProvider) SecurityKeys() (openapi.SecurityKeys, error) {

	batchKey, err := p.Cmd.Flags().GetString("batch-key")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	batchName, err := p.Cmd.Flags().GetString("batch-name")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	storageKey, err := p.Cmd.Flags().GetString("storage-key")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	storageName, err := p.Cmd.Flags().GetString("storage-name")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	k := openapi.AzureSecurityKeys{
		BatchKey:    batchKey,
		BatchName:   batchName,
		StorageKey:  storageKey,
		StorageName: storageName,
	}

	return openapi.SecurityKeys{
		Discriminator: string(AZURE),
		BatchKey:      k.BatchKey,
		BatchName:     k.BatchName,
		StorageKey:    k.StorageKey,
		StorageName:   k.StorageName,
	}, nil
}
