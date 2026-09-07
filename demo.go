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
	_, port, err := net.SplitHostPort(addr)
	if err != nil || port == "" {
		req.Headers().Set("X-Real-Port", "0")
		return true, 0
	}

	req.Headers().Set("X-Real-Port", port)
	return true, 0
}
