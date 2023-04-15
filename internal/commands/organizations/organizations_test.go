package organizations_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"tower-cli-go/internal/commands"
	tu "tower-cli-go/internal/utils/testutils"

	"github.com/stretchr/testify/assert"
)

var testDataDir = "../testdata"

func TestList(t *testing.T) {

	tRes := tu.TestResourceFolder(testDataDir)

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
				{"GET", "/user-info", tRes("user.json")},
				{"GET", "/user/1264/workspaces", tRes("workspaces/workspaces_list.json")},
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
				{"GET", "/user-info", tRes("user.json")},
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
		cmd := commands.NewRootCmd()
		cmd.SetOut(b)
		cmd.SetErr(b)
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

func TestView(t *testing.T) {

	tRes := tu.TestResourceFolder(testDataDir)

	testCases := []struct {
		epMocks []struct {
			method  string
			path    string
			jsonRes string
		}
		args        []string
		outChecker func(t *testing.T, out string)
	}{
		{
			epMocks: []struct{ method string; path string; jsonRes string; }{
				{"GET", "/user-info", tRes("user.json")},
				{"GET", "/user/1264/workspaces", tRes("workspaces/workspaces_list.json")},
				{"GET", "/orgs/27736513644467", tRes("organizations/organizations_view.json")},
			},
			args: []string{
				"organizations", "view", "-n", "organization1", 
				"-o", "json", "-t", "<tower_token>",
			},
			outChecker: func(t *testing.T, out string) { assert.JSONEq(t,
				`{
					"organization": {
						"orgId": 27736513644467,
						"name": "organization1",
						"fullName": "organization1"
					}
				}`, out)
			},
		},
		{
			epMocks: []struct{ method string; path string; jsonRes string; }{
				{"GET", "/user-info", tRes("user.json")},
				{"GET", "/user/1264/workspaces", tRes("workspaces/workspaces_list.json")},
				{"GET", "/orgs/27736513644467", "{}"},
			},
			args: []string{
				"organizations", "view", "-n", "organization1", 
				"-o", "json", "-t", "<tower_token>",
			},
			outChecker: func(t *testing.T, out string) { assert.JSONEq(t,
				`{
					"organization": {}
				}`, out)
			},
		},
		{
			epMocks: []struct{ method string; path string; jsonRes string; }{
				{"GET", "/user-info", tRes("user.json")},
				{"GET", "/user/1264/workspaces", tRes("workspaces/workspaces_list.json")},
				{"GET", "/orgs/27736513644467", tRes("organizations/organizations_view.json")},
			},
			args: []string{
				"organizations", "view", // no ID/name provided
				"-o", "json", "-t", "<tower_token>",
			},
			outChecker: func(t *testing.T, actual string) {
				assert.Contains(t, actual, "organization ID or name needed")
			},
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
		cmd := commands.NewRootCmd()
		cmd.SetOut(b)
		cmd.SetErr(b)
		cmd.SetArgs(append(tc.args, "--url", server.URL))
		cmd.Execute()

		// check cmd output
		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}
		outStr := string(out)

		// assertions
		tc.outChecker(t, outStr)

		// cleanup before next test case
		server.Close()
	}

}
