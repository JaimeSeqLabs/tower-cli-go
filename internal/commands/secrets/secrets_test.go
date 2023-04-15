package secrets_test

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

func TestAdd(t *testing.T) {

	tRes := tu.TestResourceFolder(testDataDir)

	handlers := tu.ServerMockMap{
		"/user-info": {
			"GET": tu.SendJSONResFile(tRes, "user.json"),
		},
		"/user/1264/workspaces": {
			"GET": tu.SendJSONResFile(tRes, "workspaces/workspaces_list.json"),
		},
		"/pipeline-secrets": {
			"POST": func(w http.ResponseWriter, r *http.Request) {
				if tu.BodyJSONEq(r, "{\"name\":\"name03\",\"value\":\"value03\"}") {
					tu.SendJSON(w, "{\"secretId\":164410928765888}")
				}
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
		"secrets", "add", "--name", "name03", "--value", "value03",
		"-t", "<tower_token>", "--url", server.URL,
	})
	cmd.Execute()

	// check cmd output
	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	outStr := string(out)

	expected := "New secret 'name03' (164410928765888) added at user workspace"
	assert.Contains(t, outStr, expected)

	// setup cmd, passing wsp ref
	b = bytes.NewBufferString("")
	cmd = commands.NewRootCmd()
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs([]string{
		"secrets", "add", "--name", "name03", "--value", "value03",
		"-w", "organization1/workspace1",
		"-t", "<tower_token>", "--url", server.URL,
	})
	cmd.Execute()

	// check cmd output
	out, err = io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	outStr = string(out)

	expected = "New secret 'name03' (164410928765888) added at [organization1 / workspace1] workspace"
	assert.Contains(t, outStr, expected)

	// cleanup
	server.Close()
}

func TestList(t *testing.T) {

	tRes := tu.TestResourceFolder(testDataDir)

	handlers := tu.ServerMockMap{
		"/user-info": {
			"GET": tu.SendJSONResFile(tRes, "user.json"),
		},
		"/user/1264/workspaces": {
			"GET": tu.SendJSONResFile(tRes, "workspaces/workspaces_list.json"),
		},
		"/pipeline-secrets": {
			"GET": func(w http.ResponseWriter, r *http.Request) {
				tu.SendJSON(w, `
				{
					"pipelineSecrets":[
						{
							"id":5002114781502,
							"name":"name01",
							"lastUsed":null,
							"dateCreated":"2022-10-25T12:42:21Z",
							"lastUpdated":"2022-10-25T12:42:21Z"
						},
						{
							"id":171740984431657,
							"name":"name02",
							"lastUsed":null,
							"dateCreated":"2022-10-25T13:21:15Z",
							"lastUpdated":"2022-10-25T13:21:15Z"
						}
					],
					"totalSize":2
				}
				`)
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
		"secrets", "list",
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

	expected := `
	{
		"workspaceRef": "user",
		"secrets": [
			{
				"id":5002114781502,
				"name": "name01",
				"dateCreated":"2022-10-25T12:42:21Z",
				"lastUpdated":"2022-10-25T12:42:21Z",
				"lastUsed": "0001-01-01T00:00:00Z"
			},
			{
				"id":171740984431657,
				"name": "name02",
				"dateCreated":"2022-10-25T13:21:15Z",
				"lastUpdated":"2022-10-25T13:21:15Z",
				"lastUsed": "0001-01-01T00:00:00Z"
			}
		]

	}`
	assert.JSONEq(t, expected, outStr)

	// setup cmd, passing wsp ref
	b = bytes.NewBufferString("")
	cmd = commands.NewRootCmd()
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs([]string{
		"secrets", "list", "-w", "organization1/workspace1",
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

	expected = `
	{
		"workspaceRef": "[organization1 / workspace1]",
		"secrets": [
			{
				"id":5002114781502,
				"name": "name01",
				"dateCreated":"2022-10-25T12:42:21Z",
				"lastUpdated":"2022-10-25T12:42:21Z",
				"lastUsed": "0001-01-01T00:00:00Z"
			},
			{
				"id":171740984431657,
				"name": "name02",
				"dateCreated":"2022-10-25T13:21:15Z",
				"lastUpdated":"2022-10-25T13:21:15Z",
				"lastUsed": "0001-01-01T00:00:00Z"
			}
		]

	}`
	assert.JSONEq(t, expected, outStr)

	// cleanup
	server.Close()
}

func TestView(t *testing.T) {

	tRes := tu.TestResourceFolder(testDataDir)

	handlers := tu.ServerMockMap{
		"/user-info": {
			"GET": tu.SendJSONResFile(tRes, "user.json"),
		},
		"/user/1264/workspaces": {
			"GET": tu.SendJSONResFile(tRes, "workspaces/workspaces_list.json"),
		},
		"/pipeline-secrets": {
			"GET": func(w http.ResponseWriter, r *http.Request) {
				tu.SendJSON(w, `
				{
					"pipelineSecrets":[
						{
							"id":5002114781502,
							"name":"name01",
							"lastUsed":null,
							"dateCreated":"2022-10-25T12:42:21Z",
							"lastUpdated":"2022-10-25T12:42:21Z"
						},
						{
							"id":171740984431657,
							"name":"name02",
							"lastUsed":null,
							"dateCreated":"2022-10-25T13:21:15Z",
							"lastUpdated":"2022-10-25T13:21:15Z"
						}
					],
					"totalSize":2
				}
				`)
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
		"secrets", "view", "--name", "name02",
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

	expected := `
	{
		"workspaceRef": "user",
		"secret": {
			"id": 171740984431657,
			"name":"name02",
			"lastUsed":"0001-01-01T00:00:00Z",
			"dateCreated":"2022-10-25T13:21:15Z",
			"lastUpdated":"2022-10-25T13:21:15Z"
		}
	}`
	assert.JSONEq(t, expected, outStr)

	// setup cmd, passing wsp ref
	b = bytes.NewBufferString("")
	cmd = commands.NewRootCmd()
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs([]string{
		"secrets", "view", "--name", "name02",
		"-w", "organization1/workspace1",
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

	expected = `
	{
		"workspaceRef": "[organization1 / workspace1]",
		"secret": {
			"id": 171740984431657,
			"name":"name02",
			"lastUsed":"0001-01-01T00:00:00Z",
			"dateCreated":"2022-10-25T13:21:15Z",
			"lastUpdated":"2022-10-25T13:21:15Z"
		}
	}`
	assert.JSONEq(t, expected, outStr)

	// cleanup
	server.Close()
}

func TestUpdate(t *testing.T) {

	tRes := tu.TestResourceFolder(testDataDir)

	handlers := tu.ServerMockMap{
		"/user-info": {
			"GET": tu.SendJSONResFile(tRes, "user.json"),
		},
		"/user/1264/workspaces": {
			"GET": tu.SendJSONResFile(tRes, "workspaces/workspaces_list.json"),
		},
		"/pipeline-secrets": {
			"GET": func(w http.ResponseWriter, r *http.Request) {
				tu.SendJSON(w, `
				{
					"pipelineSecrets":[
						{
							"id":5002114781502,
							"name":"name01",
							"lastUsed":null,
							"dateCreated":"2022-10-25T12:42:21Z",
							"lastUpdated":"2022-10-25T12:42:21Z"
						},
						{
							"id":171740984431657,
							"name":"name02",
							"lastUsed":null,
							"dateCreated":"2022-10-25T13:21:15Z",
							"lastUpdated":"2022-10-25T13:21:15Z"
						}
					],
					"totalSize":2
				}
				`)
			},
		},
		"/pipeline-secrets/171740984431657": {
			"PUT": func(w http.ResponseWriter, r *http.Request) {
				if tu.BodyJSONEq(r, `{"value": "updateValue"}`) {
					w.WriteHeader(204)
				}
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
		"secrets", "update", "--name", "name02", "--value", "updateValue",
		"-t", "<tower_token>", "--url", server.URL,
	})
	cmd.Execute()

	// check cmd output
	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	outStr := string(out)

	expected := "Secret 'name02' updated at user workspace"
	assert.Contains(t, outStr, expected)

	// setup cmd, passing wsp ref
	b = bytes.NewBufferString("")
	cmd = commands.NewRootCmd()
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs([]string{
		"secrets", "update", "--name", "name02", "--value", "updateValue",
		"-w", "organization1/workspace1",
		"-t", "<tower_token>", "--url", server.URL,
	})
	cmd.Execute()

	// check cmd output
	out, err = io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	outStr = string(out)

	expected = "Secret 'name02' updated at [organization1 / workspace1] workspace"
	assert.Contains(t, outStr, expected)

	// cleanup
	server.Close()
}

func TestDelete(t *testing.T) {

	tRes := tu.TestResourceFolder(testDataDir)

	handlers := tu.ServerMockMap{
		"/user-info": {
			"GET": tu.SendJSONResFile(tRes, "user.json"),
		},
		"/user/1264/workspaces": {
			"GET": tu.SendJSONResFile(tRes, "workspaces/workspaces_list.json"),
		},
		"/pipeline-secrets": {
			"GET": func(w http.ResponseWriter, r *http.Request) {
				tu.SendJSON(w, `
				{
					"pipelineSecrets":[
						{
							"id":5002114781502,
							"name":"name01",
							"lastUsed":null,
							"dateCreated":"2022-10-25T12:42:21Z",
							"lastUpdated":"2022-10-25T12:42:21Z"
						},
						{
							"id":171740984431657,
							"name":"name02",
							"lastUsed":null,
							"dateCreated":"2022-10-25T13:21:15Z",
							"lastUpdated":"2022-10-25T13:21:15Z"
						}
					],
					"totalSize":2
				}
				`)
			},
		},
		"/pipeline-secrets/164410928765888": {
			"DELETE": func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(204) },
		},
	}

	server := tu.MockServer(handlers)

	// setup cmd, no wsp ref
	b := bytes.NewBufferString("")
	cmd := commands.NewRootCmd()
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs([]string{
		"secrets", "delete", "--name", "name02",
		"-t", "<tower_token>", "--url", server.URL,
	})
	cmd.Execute()

	// check cmd output
	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	outStr := string(out)

	expected := "Secret 'name02' deleted at user workspace"
	assert.Contains(t, outStr, expected)

	// setup cmd, passing wsp ref
	b = bytes.NewBufferString("")
	cmd = commands.NewRootCmd()
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs([]string{
		"secrets", "delete", "--name", "name02",
		"-w", "organization1/workspace1",
		"-t", "<tower_token>", "--url", server.URL,
	})
	cmd.Execute()

	// check cmd output
	out, err = io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	outStr = string(out)

	expected = "Secret 'name02' deleted at [organization1 / workspace1] workspace"
	assert.Contains(t, outStr, expected)

	// cleanup
	server.Close()
}
