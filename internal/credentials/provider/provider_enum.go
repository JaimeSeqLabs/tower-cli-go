package provider

type CredentialsProviderEnum string

const (
	AWS           CredentialsProviderEnum = "aws"
	AZURE         CredentialsProviderEnum = "azure"
	GOOGLE        CredentialsProviderEnum = "google"
	GITHUB        CredentialsProviderEnum = "github"
	GITLAB        CredentialsProviderEnum = "gitlab"
	BITBUCKET     CredentialsProviderEnum = "bitbucket"
	SSH           CredentialsProviderEnum = "ssh"
	K8S           CredentialsProviderEnum = "k8s"
	CONTAINER_REG CredentialsProviderEnum = "container-reg"
	TW_AGENT      CredentialsProviderEnum = "tw-agent"
	CODECOMMIT    CredentialsProviderEnum = "codecommit"
)

func EnumValues() []CredentialsProviderEnum {
	return []CredentialsProviderEnum{
		AWS,
		AZURE,
		GOOGLE,
		GITHUB,
		GITLAB,
		BITBUCKET,
		SSH,
		K8S,
		CONTAINER_REG,
		TW_AGENT,
		CODECOMMIT,
	}
}

func FromValue(v string) CredentialsProviderEnum {
	for _, e := range EnumValues() {
		if string(e) == v {
			return e
		}
	}
	return CredentialsProviderEnum("")
}
