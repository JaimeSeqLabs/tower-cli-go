package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"tower-cli-go/internal"
	"tower-cli-go/internal/formatters"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

func NewInfoCmd() *cobra.Command {
	
	infoCmd := &cobra.Command{
		Use:   "info",
		Short: "System info and health status",
		RunE: func(cmd *cobra.Command, args []string) error {
	
			wrapper := internal.NewApiFor(cmd)

			var connectionCheck int = 1
			var versionCheck int = -1
			var credentialsCheck int = -1

			// TODO: use go:embed
			var usrName string
			var towerVersion string
			var towerApiVersion string
			var cliVersion string = "0.0.1 (db6b6487568a874437861d8b652d440cad5c9c90)"
			var cliApiVersion string = "0.0.1"

			infoResponse, _, err := wrapper.Api.Info(wrapper.Ctx)
			if err != nil {
				connectionCheck = 0
			}

			if infoResponse.ServiceInfo.ApiVersion != "" {
				towerVersion = infoResponse.ServiceInfo.Version
				towerApiVersion = infoResponse.ServiceInfo.ApiVersion
				// if tower version lower than cli's fail version check
				if comp := strings.Compare(towerApiVersion, cliApiVersion); comp >= 0 {
					versionCheck = 1
				} else {
					versionCheck =  0
				}
			}

			if connectionCheck == 1 {

				userDesc, httpRes, err := wrapper.Api.Profile(wrapper.Ctx)
				if err != nil {
					if httpRes.StatusCode == 401 {
						credentialsCheck = 0
					}
				} else {
					usrName = userDesc.User.UserName
					credentialsCheck = 1
				}
			
			}

			result := InfoResponse {
				connectionCheck,
				versionCheck,
				credentialsCheck,
				make(map[string]string),
			}

			result.setOpt("cliApiVersion", cliApiVersion)
			result.setOpt("cliVersion", cliVersion)
			result.setOpt("towerApiVersion", towerApiVersion)
			result.setOpt("towerApiEndpoint", wrapper.ApiUrl())
			result.setOpt("towerVersion", towerVersion)
			result.setOpt("userName", usrName)

			return formatters.PrintTo(cmd.OutOrStdout(), result)
		},
	}

	return infoCmd
}

type InfoResponse struct {
	ConnectionCheck int `json:"connectionCheck"`
	VersionCheck int `json:"versionCheck"`
	CredentialsCheck int `json:"credentialsCheck"`
	Opts map[string]string `json:"opts"`
}

func (info InfoResponse) getOpt(key, defaultValue string) string {
	if val, found := info.Opts[key]; found {
		return val
	} else {
		return defaultValue
	}
}

func (info InfoResponse) setOpt(key, value string) {
	if value != "" {
		info.Opts[key] = value
	}
}

func (info InfoResponse) WriteAsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(info)
}

func (info InfoResponse) WriteAsTable(w io.Writer) error {

	text.EnableColors()
	defer text.DisableColors()

	createTable := func() table.Writer {
		t := table.NewWriter()
		t.SetOutputMirror(w)
		t.SetStyle(table.StyleLight)
		t.Style().Options.SeparateRows = true
		t.Style().Format.Header = text.FormatDefault
		return t
	}
	
	fmt.Fprintf(w, "\nDetails\n\n")

	undef := colorUndef("UNDEFINED")
	
	t := createTable()
	t.AppendRows([]table.Row{
		{ "Tower API endpoint", 		info.getOpt("towerApiEndpoint", undef) },
		{ "Tower API version", 			info.getOpt("towerApiVersion", undef) },
		{ "Tower version", 				info.getOpt("towerVersion", undef) },
		{ "CLI version", 				info.getOpt("cliVersion", undef) },
		{ "CLI minimum API version", 	info.getOpt("cliApiVersion", undef) },
		{ "Authenticated user", 		info.getOpt("userName", undef) },		
	})
	t.Render()

	fmt.Fprintf(w, "\nSystem health status\n\n")

	t = createTable()
	t.AppendRows([]table.Row{
		{ "Remote API server connection check", toStr(info.ConnectionCheck) },
        { "Tower API version check", toStr(info.VersionCheck) },
        { "Authentication API credential's token", toStr(info.CredentialsCheck) },
	})
	t.Render()

	fmt.Fprintln(w)

	if info.ConnectionCheck == 0 {
		if ep := info.Opts["towerApiEndpoint"]; strings.Contains(ep, "/api") {
			fmt.Fprintf(w, colorFail("Tower API URL %s it is not available")+"\n\n", ep)
		} else {
			fmt.Fprintf(w, colorFail("Tower API URL %s it is not available (did you mean %s/api?)")+"\n\n", ep, ep)
		}
	}

	if info.VersionCheck == 0 {
		fmt.Fprintf(w, 
			colorFail("Tower API version is %s while the minimum required version to be fully compatible is %s")+"\n\n",
			info.Opts["towerApiVersion"],
			info.Opts["cliApiVersion"],
		)
	}

	if info.CredentialsCheck == 0 {
		fmt.Fprintln(w, colorFail("Please check that your current access token is valid and active")+"\n\n")
	}
	
	if info.ConnectionCheck + info.VersionCheck + info.CredentialsCheck != 3 {
		return fmt.Errorf("health checks failed")
	}

	return nil
}

func toStr(b int) string {
	if b == 0 {
		return colorFail("FAILED")
	}
	if b == 1 {
		return colorOk("OK")
	}
	return colorUndef("UNDEFINED")
}

func colorOk(s string) string {
	return text.Colors{text.FgGreen}.Sprint(s) 
}

func colorFail(s string) string {
	return text.Colors{text.FgHiRed}.Sprint(s) 
}

func colorUndef(s string) string {
	return text.Colors{text.FgYellow}.Sprint(s) 
}


