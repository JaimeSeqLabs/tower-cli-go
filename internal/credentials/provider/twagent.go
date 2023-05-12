package provider

import (
	"github.com/spf13/cobra"
	"openapi"
)

func NewTwAgentCredentialsProvider(cmd *cobra.Command) CredentialsProvider {
	return twAgentCredsProvider{cmd}
}

type twAgentCredsProvider struct {
	Cmd *cobra.Command
}

func (p twAgentCredsProvider) BaseUrl() string {
	return ""
}

func (p twAgentCredsProvider) Type() CredentialsProviderEnum {
	return TW_AGENT
}

func (p twAgentCredsProvider) SecurityKeys() (openapi.SecurityKeys, error) {

	connectionID, err := p.Cmd.Flags().GetString("connection-id")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	workDir, err := p.Cmd.Flags().GetString("work-dir")
	if err != nil {
		return openapi.SecurityKeys{}, err
	}

	k := openapi.AgentSecurityKeys{
		ConnectionId: connectionID,
		WorkDir:      workDir,
	}

	return openapi.SecurityKeys{
		Discriminator: string(TW_AGENT),
		ConnectionId:  k.ConnectionId, // note 'd' lowercase
		WorkDir:       k.WorkDir,
	}, nil
}
