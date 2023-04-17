package main

import (
	"bytes"
	//"fmt"
	"io"
	"strings"
	"syscall/js"
	"tower-cli-go/internal/commands"
)

const (
	exportedFuncName = "TowerCli"
)

func main() {

	jsCallCmd := createAsyncCallFunc()
	//fmt.Println("> Function created")

	// export under global namespace in js
	js.Global().Set(exportedFuncName, jsCallCmd)
	//fmt.Println("> Global function exported")

	// to prevent wasm from collecting cli's memory context we must not exit the main function
	// this parks the routine forever (or until wasm context gets explicitly destroyed)
	//fmt.Println("> Parking main routine")
	<-make(chan any)
}

// async function TowerCli(params: string): {stdout: string, stderr: string}
func createAsyncCallFunc() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {

		// create a Promise handle
		handler := js.FuncOf(func(this js.Value, resolveRejectArgs []js.Value) any {

			resolve := resolveRejectArgs[0]
			reject := resolveRejectArgs[1]

			// launch separate goroutine so js loop doesn't block
			go func() {

				_cliParams := args[0]
				if _cliParams.Type() != js.TypeString {
					reject.Invoke("cli parameters must be a string")
					return
				}
				cliParams := _cliParams.String()

				outBuffer := bytes.NewBufferString("")
				errBuffer := bytes.NewBufferString("")

				root := commands.NewRootCmd()
				root.SetArgs(strings.Split(cliParams, " "))
				root.SetOut(outBuffer)
				root.SetErr(errBuffer)
				root.Execute()
				
				outRes, err := io.ReadAll(outBuffer)
				if err != nil {
					reject.Invoke(err.Error())
					return
				}
				outStr := string(outRes)
				
				errRes, err := io.ReadAll(errBuffer)
				if err != nil {
					reject.Invoke(err.Error())
					return
				}
				errStr := string(errRes)

				result := map[string] any {
					"stdout": outStr,
					"stderr": errStr,
				}
				resolve.Invoke(result)

			}()

			// handler of a promise never returns any value
			return nil
		})

		promiseCtor := js.Global().Get("Promise")
		return promiseCtor.New(handler)
	})
}
