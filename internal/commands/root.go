package commands

import (
	"fmt"
	"os"
	"tower-cli-go/internal/commands/organizations"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile      string
	accessToken  string
	serverUrl    string
	outputFormat string
	verbose      bool
	insecure     bool
	showVersion  bool
)

var RootCmd = &cobra.Command{
	Use:   "tw",
	Short: "Tower command line tool",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		if show, _ := cmd.Flags().GetBool("version"); show {
			fmt.Println("TODO: show version") // TODO
			os.Exit(0)
		}
	
	},
	Run: func(cmd *cobra.Command, args []string) {
		// nothing
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	// load cfg from env/file
	cobra.OnInitialize(initViperCfg)

	// persistent global flags, shared with subcommands
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (none by default)")
	RootCmd.PersistentFlags().StringVarP(&accessToken, "access-token", "t", "", "Tower personal access token (TOWER_ACCESS_TOKEN)")
	RootCmd.PersistentFlags().StringVarP(&serverUrl, "url", "u", "https://api.tower.nf", "Tower server API endpoint URL (TOWER_API_ENDPOINT)")
	RootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "", "Show output in defined format (only 'json' allowed)")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Show HTTP request/response logs at stderr")
	RootCmd.PersistentFlags().BoolVar(&insecure, "insecure", false, "Explicitly allow to connect to a non-SSL secured Tower server (this is not recommended)")
	RootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "V", false, "Print version information and exit")

	// bind flags to viper sources
	viper.BindPFlag("url", RootCmd.PersistentFlags().Lookup("url"))
	viper.BindEnv("url", "TOWER_API_ENDPOINT")
	viper.BindPFlag("output", RootCmd.PersistentFlags().Lookup("output"))
	viper.BindPFlag("verbose", RootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("access-token", RootCmd.PersistentFlags().Lookup("access-token"))
	viper.BindEnv("access-token", "TOWER_ACCESS_TOKEN")
	viper.BindPFlag("insecure", RootCmd.PersistentFlags().Lookup("insecure"))

	// add subcommands
	RootCmd.AddCommand(
		InfoCmd,
		organizations.OrganizationsCmd,
	)

}

func initViperCfg() {

	// load from file if any
	if cfgFile != "" {

		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			cobra.CheckErr(err) // exits
		}
	}

	// load cfg from matching environment vars
	viper.SetEnvPrefix("TW_")
	viper.AutomaticEnv()

}
