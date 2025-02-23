GOOS=js GOARCH=wasm tinygo build -o production.wasm -no-debug -scheduler=none -panic=trap -gc=leaking
