package router

import (
	"github.com/valyala/fasthttp"
)

// PassFlag indicates whether Handlers.Handler and Handlers.PostHandler are to be skipped.
type PassFlag int

// Flags for Handlers.Checkpoint
const (
	Continue PassFlag = iota
	Stop
)

// CheckHandler is like fasthttp.RequestHandler but returning a flag.
type CheckHandler func(ctx *fasthttp.RequestCtx) PassFlag

// Handlers is a collection of request handlers.
type Handlers struct {
	Handler fasthttp.RequestHandler
	Wrapper *Wrapper
}

// SetHandler return a copy of Handlers with its Handler are set and
// Wrapper point to the same Wrapper.
func (h *Handlers) SetHandler(r fasthttp.RequestHandler) *Handlers {
	c := &Handlers{}
	c.Handler = r
	c.Wrapper = h.Wrapper
	return c
}
