build:
	GOOS=js GOARCH=wasm go build -o assets/main.wasm cmd/wasm/main.go 
	go build cmd/server/main.go 