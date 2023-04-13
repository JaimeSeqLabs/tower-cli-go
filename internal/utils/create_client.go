package utils

import (
	"fmt"
	"openapi"

	"github.com/spf13/viper"
)

func GenerateClientFromCfg() *openapi.APIClient {

	cfg := openapi.NewConfiguration()

	cfg.Debug = viper.GetBool("verbose")
	cfg.BasePath = viper.GetString("url")
	if token := viper.GetString("access-token"); token != "" {
		cfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	} else {
		panic("missing tower access token")
	}
	// TODO: if insecure set create a custom insecure http client

	api := openapi.NewAPIClient(cfg)
	if api == nil {
		panic("unable to create an openapi client with current cfg")
	}

	return api
}