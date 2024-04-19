GOROOT_PATH = $(shell go env GOROOT)
init:
	cp $(GOROOT_PATH)/misc/wasm/wasm_exec.js assets

build:
	GOOS=js GOARCH=wasm go build -o assets/main.wasm cmd/wasm/main.go 
	go build -o ConwayWASM cmd/server/main.go 