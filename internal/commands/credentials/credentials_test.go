package credentials_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	"tower-cli-go/internal/commands"
	tu "tower-cli-go/internal/utils/testutils"

	"github.com/stretchr/testify/assert"
)

var testDataDir = "../testdata"

func TestList(t *testing.T) {

	tRes := tu.TestResourceFolder(testDataDir)

	handlers := tu.ServerMockMap {
		"/user-info": {
			"GET": tu.SendJSONResFile(tRes, "user.json"),
		},
		"/user/1264/workspaces": {
			"GET": tu.SendJSONResFile(tRes, "workspaces/workspaces_list.json"),
		},
		"/credentials": {
			"GET": tu.SendJSONResFile(tRes, "credentials_list.json"),
		},
	}

	server := tu.MockServer(handlers)

	// setup cmd, no wsp ref
	b := bytes.NewBufferString("")
	cmd := commands.NewRootCmd()
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs([]string{
		"credentials", "list",
		"-o", "json",
		"-t", "<tower_token>", "--url", server.URL,
	})
	cmd.Execute()

	// check cmd output
	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	outStr := string(out)

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
	}`, outStr)

	// setup cmd
	b = bytes.NewBufferString("")
	cmd = commands.NewRootCmd()
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs([]string{
		"credentials", "list",
		"-w", "organization2/workspace2",
		"-o", "json",
		"-t", "<tower_token>", "--url", server.URL,
	})
	cmd.Execute()

	// check cmd output
	out, err = io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	outStr = string(out)

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
	}`, outStr)

	// empty response case
	handlers["/credentials"] = map[string]http.HandlerFunc {
			"GET": func(w http.ResponseWriter, r *http.Request) {
				tu.SendJSON(w, "{}")
			},
	}

	// setup cmd
	b = bytes.NewBufferString("")
	cmd = commands.NewRootCmd()
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs([]string{
		"credentials", "list",
		"-w", "organization2/workspace2",
		"-o", "json",
		"-t", "<tower_token>", "--url", server.URL,
	})
	cmd.Execute()

	// check cmd output
	out, err = io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	outStr = string(out)

	assert.JSONEq(t,
		`{
		"workspaceRef": "[organization2 / workspace2]",
		"credentials": []
	}`, outStr)

	// cleanup
	server.Close()
}

func TestDelete(t *testing.T) {

	credID := "1cz5A8cuBkB5iJliCwJCFU"
	noExistCredID := "2d06B9dvCkB5iJliCwJDGV"
	unauthCredID := "3e17C0ewDmC5iJliCwKEHX"

	handlers := tu.ServerMockMap {
		"/credentials/"+credID: {
			"DELETE": func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(204)
			},
		},
		"/credentials/"+noExistCredID: {
			"DELETE": func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(403)
			},
		},
		"/credentials/"+unauthCredID: {
			"DELETE": func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(401)
			},
		},
	}

	server := tu.MockServer(handlers)

	// setup cmd, no wsp ref
	b := bytes.NewBufferString("")
	cmd := commands.NewRootCmd()
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs([]string{
		"credentials", "delete",
		"-i", credID,
		"-o", "json",
		"-t", "<tower_token>", "--url", server.URL,
	})
	cmd.Execute()

	// check cmd output
	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	outStr := string(out)

	assert.Contains(t, outStr, credID, "user") // user wsp

	// setup cmd, non existing id
	b = bytes.NewBufferString("")
	cmd = commands.NewRootCmd()
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs([]string{
		"credentials", "delete",
		"-i", unauthCredID,
		"-o", "json",
		"-t", "<tower_token>", "--url", server.URL,
	})
	cmd.Execute()

	// check cmd output
	out, err = io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	outStr = string(out)

	assert.Contains(t, outStr, "Unauthorized", "user") // user wsp

}
	