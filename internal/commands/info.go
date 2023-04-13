package commands

import (
	"context"
	"fmt"
	"tower-cli-go/internal/utils"

	"github.com/spf13/cobra"
)

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get API info",
	Run: func(cmd *cobra.Command, args []string) {

		api := utils.GenerateClientFromCfg()

		response, _, err := api.DefaultApi.Info(context.TODO())
		if err != nil {
			cobra.CheckErr(err)
		}

		apiVersion := response.ServiceInfo.ApiVersion

		fmt.Printf("API Version is %s\n", apiVersion)
	},
}
