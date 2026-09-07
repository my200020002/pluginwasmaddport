package main

import (
	"net"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleRequestFn = handleRequest
}

func handleRequest(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	req.Headers().Set("X-Wasm-Debug", "triggered")

	addr := req.GetSourceAddr()
	
	req.Headers().Set("X-Debug-Source-Addr", addr)


	return true, 0
}
