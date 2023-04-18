package provider

import (
	"openapi"
)

type CredentialsProvider interface {
	BaseUrl() string
	Type() CredentialsProviderEnum
	SecurityKeys() (openapi.SecurityKeys, error)
}
