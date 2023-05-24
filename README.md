# Nextflow Tower CLI - Go version


Implementation of [tower-cli](https://github.com/seqeralabs/tower-cli) in Go.

> Work in progress

## WASM Demo
This is a separate branch showcasing `tower-cli-go` compiled to WASM running in the browser.

How to run the demo:

1) Build Tower Go sdk:
```shell
sh gen_tower_sdk.sh
```

2) Compile `tower-cli-go` to WASM:
```shell
sh build_wasm.sh
```

3) Run an instance of Tower in localhost (local instance is needed because of CORS)

4) Serve the demo page:
```shell
sh run_server.sh
```

5) Go to http://localhost:1313 to use the cli


## Development

### Required dependencies:
- Go 1.19+
- For generating client from OpenAPI spec:
    - java 11
    - Gradle 6.6.1+ (optional, local gradlew script can be used)
    - `sed` command in shell

### Cloning sources
```
git clone --recurse-submodules https://github.com/JaimeSeqLabs/tower-cli-go.git
```

### Generating Tower go sdk
> Java 11 is required

In project root run:
```
go generate ./...
```
Or run:
```
sh gen_tower_sdk.sh
```
> These commands perform some patches to the generated code using `sed`

### Building CLI
```
go build -o ./tw cmd/cli/main.go
```
You can also run the CLI without building the binary like this:
```
go run cmd/cli/main.go --help
```

### Usage example
```
$ export TOWER_ACCESS_TOKEN=<your_tower_token>
$ ./tw organizations list
$ go run cmd/cli/main.go organizations list
```

## Status

Actions:
- [ ] Add
- [ ] View
- [ ] List
- [ ] Update
- [ ] Delete

Collaborators:
- [ ] List

Compute-envs:
- [ ] Add
- [ ] View
- [ ] List
- [ ] Update
- [ ] Delete
- [ ] Export
- [ ] Import
- [ ] Primary

Credentials:
- [X] Add
- [X] List
- [X] Update
- [X] Delete

Datasets:
- [ ] Add
- [ ] List
- [ ] View
- [ ] Update
- [ ] Delete
- [ ] Download
- [ ] Url

Info:
- [X]

Launch:
- [ ]

Members:
- [ ] Add
- [ ] List
- [ ] Update
- [ ] Delete
- [ ] Leave

Organizations:
- [X] Add
- [X] View
- [X] List
- [X] Update
- [X] Delete

Participants:
- [ ] Add
- [ ] List
- [ ] Update
- [ ] Delete
- [ ] Leave

Pipelines:
- [ ] Add
- [ ] View
- [ ] List
- [ ] Update
- [ ] Delete
- [ ] Export
- [ ] Import

Runs:
- [ ] View
- [ ] List
- [ ] Relaunch
- [ ] Cancel
- [ ] Delete

Teams:
- [ ] Add
- [ ] List
- [ ] Members
- [ ] Delete

Workspaces:
- [X] Add
- [X] List
- [X] View
- [X] Update
- [X] Delete
- [X] Leave

Secrets:
- [X] Add
- [X] View
- [X] List
- [X] Update
- [X] Delete
