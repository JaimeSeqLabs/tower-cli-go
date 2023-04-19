package credentials_test

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
		args       []string
		outChecker func(t *testing.T, out string)
	}{
		{
			epMocks: []struct {
				method  string
				path    string
				jsonRes string
			}{
				{"GET", "/user-info", tRes("user.json")},
				{"GET", "/user", tRes("user.json")},
				{"GET", "/user/1264/workspaces", tRes("workspaces/workspaces_list.json")},
				{"GET", "/credentials", tRes("credentials_list.json")},
			},
			args: []string{
				"credentials", "list",
				"-o", "json", "-t", "<tower_token>",
				// no workspace ref
			},
			outChecker: func(t *testing.T, out string) {
				assert.JSONEq(t,
					`{
					"workspaceRef": "user",
					"credentials": [
						{
							"id": "2ba2oekqeTEBzwSDgXg7xf",
							"name": "ssh",
							"provider": "ssh",
							"lastUsed": "2021-09-06T08:53:51Z",
							"dateCreated": "2021-09-06T06:54:53Z",
							"lastUpdated": "2021-09-06T06:54:53Z",
							"keys": {
								"privateKey": "privateKey",
								"passphrase": "passphrase",
								"discriminator": "ssh"
							}
						},
						{
							"id": "57Ic6reczFn78H1DTaaXkp",
							"name": "azure",
							"provider": "azure",
							"dateCreated": "2021-09-07T13:50:21Z",
							"lastUpdated": "2021-09-07T13:50:21Z",
							"lastUsed": "0001-01-01T00:00:00Z",
							"keys": {
								"batchName": "batchName",
								"batchKey": "batchKey",
								"storageName": "storageName",
								"storageKey": "storageKey",
								"discriminator": "azure"
							}
						},
						{
							"id": "5aoTHK1PcXdIldjb7EarQx",
							"name": "aws",
							"provider": "aws",
							"lastUsed": "2022-09-06T08:53:51Z",
							"dateCreated": "2022-09-06T06:54:53Z",
							"lastUpdated": "2022-09-06T06:54:53Z",
							"keys": {
								"accessKey": "accessKey",
								"secretKey": "secretKey",
								"assumeRoleArn": "awsRole",
								"discriminator": "aws"
							}
						}
					]
				}`, out)
			},
		},
		{
			epMocks: []struct {
				method  string
				path    string
				jsonRes string
			}{
				{"GET", "/user-info", tRes("user.json")},
				{"GET", "/user", tRes("user.json")},
				{"GET", "/user/1264/workspaces", tRes("workspaces/workspaces_list.json")},
				{"GET", "/credentials", tRes("credentials_list.json")},
			},
			args: []string{
				"credentials", "list",
				"-o", "json", "-t", "<tower_token>",
				"-w", "organization2/workspace2",
			},
			outChecker: func(t *testing.T, out string) {
				assert.JSONEq(t,
					`{
					"workspaceRef": "[organization2 / workspace2]",
					"credentials": [
						{
							"id": "2ba2oekqeTEBzwSDgXg7xf",
							"name": "ssh",
							"provider": "ssh",
							"lastUsed": "2021-09-06T08:53:51Z",
							"dateCreated": "2021-09-06T06:54:53Z",
							"lastUpdated": "2021-09-06T06:54:53Z",
							"keys": {
								"privateKey": "privateKey",
								"passphrase": "passphrase",
								"discriminator": "ssh"
							}
						},
						{
							"id": "57Ic6reczFn78H1DTaaXkp",
							"name": "azure",
							"provider": "azure",
							"dateCreated": "2021-09-07T13:50:21Z",
							"lastUpdated": "2021-09-07T13:50:21Z",
							"lastUsed": "0001-01-01T00:00:00Z",
							"keys": {
								"batchName": "batchName",
								"batchKey": "batchKey",
								"storageName": "storageName",
								"storageKey": "storageKey",
								"discriminator": "azure"
							}
						},
						{
							"id": "5aoTHK1PcXdIldjb7EarQx",
							"name": "aws",
							"provider": "aws",
							"lastUsed": "2022-09-06T08:53:51Z",
							"dateCreated": "2022-09-06T06:54:53Z",
							"lastUpdated": "2022-09-06T06:54:53Z",
							"keys": {
								"accessKey": "accessKey",
								"secretKey": "secretKey",
								"assumeRoleArn": "awsRole",
								"discriminator": "aws"
							}
						}
					]
				}`, out)
			},
		},
		{
			epMocks: []struct {
				method  string
				path    string
				jsonRes string
			}{
				{"GET", "/user-info", tRes("user.json")},
				{"GET", "/user", tRes("user.json")},
				{"GET", "/user/1264/workspaces", tRes("workspaces/workspaces_list.json")},
				{"GET", "/credentials", "{}"},
			},
			args: []string{
				"credentials", "list",
				"-o", "json", "-t", "<tower_token>",
				"-w", "organization2/workspace2",
			},
			outChecker: func(t *testing.T, out string) {
				assert.JSONEq(t,
					`{
					"workspaceRef": "[organization2 / workspace2]",
					"credentials": []
				}`, out)
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
