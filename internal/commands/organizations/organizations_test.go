package organizations_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"tower-cli-go/internal/commands"

	"github.com/stretchr/testify/assert"
)

var testDataDir = "../testdata"

func getTestResource(path string) string {
	path = filepath.Join(testDataDir, path)
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func TestList(t *testing.T) {

	testCases := []struct {
		epMocks []struct {
			method  string
			path    string
			jsonRes string
		}
		args        []string
		expectedOut string
	}{
		{
			epMocks: []struct{ method string; path string; jsonRes string; }{
				{"GET", "/user-info", getTestResource("user.json")},
				{"GET", "/user/1264/workspaces", getTestResource("workspaces/workspaces_list.json")},
			},
			args: []string{
				"organizations", "list", "-o", "json", "-t", "<tower_token>",
			},
			expectedOut: `
			{
				"userName": "jordi",
				"organizations": [
					{
						"orgId": 27736513644467,
						"orgName": "organization1"
					},
					{
						"orgId": 37736513644467,
						"orgName": "organization2",
						"orgLogoUrl": "https://tower.nf/api/avatars/4ZB6NbWph5oBIDDICKjDEJ"
					}
				]
			}`,
		},
		{
			epMocks: []struct{ method string; path string; jsonRes string; }{
				{"GET", "/user-info", getTestResource("user.json")},
				{"GET", "/user/1264/workspaces", "{}"},
			},
			args: []string{
				"organizations", "list", "-o", "json", "-t", "<tower_token>",
			},
			expectedOut: `{
				"userName": "jordi",
				"organizations": []
			}`,
		},
	}

	for _, tc := range testCases {

		// setup fake endpoints
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, ep := range tc.epMocks {

				if ep.path == r.URL.Path && ep.method == r.Method {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(200)
					_, _ = w.Write([]byte(ep.jsonRes))
				}

			}
		}))

		// setup cmd
		b := bytes.NewBufferString("")
		cmd := commands.RootCmd
		cmd.SetOut(b)
		cmd.SetArgs(append(tc.args, "--url", server.URL))
		cmd.Execute()

		// check cmd output
		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}
		outStr := string(out)

		assert.JSONEq(t, tc.expectedOut, outStr)

		// cleanup before next test case
		server.Close()
	}

}
