export GOARCH=wasm
export GOOS=js
split.wasm:*.go
	go build -o split.wasm -ldflags "-s -w"
	cp split.wasm ../static/split.wasm
	cp $(go env GOROOT)/misc/wasm/wasm_exec.js ../static/


