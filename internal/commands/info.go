package commands

import (
	"context"
	"fmt"
	"openapi"

	"github.com/spf13/cobra"
)	

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get API info",
	Run: func(cmd *cobra.Command, args []string) {
		
		cfg := openapi.NewConfiguration()
		cfg.Debug = true
		cfg.BasePath = "https://api.tower.nf"

		api := openapi.NewAPIClient(cfg)

		response, _, err := api.DefaultApi.Info(context.TODO())
		if err != nil {
			panic(err)
		}

		apiVersion := response.ServiceInfo.ApiVersion

		fmt.Printf("API Version is %s\n", apiVersion)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
