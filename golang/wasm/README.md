# HOWTO

GOOS=js GOARCH=wasm go build
goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
