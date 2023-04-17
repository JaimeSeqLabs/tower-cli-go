
GOOS=js GOARCH=wasm go build -o static/main.wasm cmd/wasm/main_js_wasm.go

cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./static
