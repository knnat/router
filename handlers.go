package router

import (
	"github.com/valyala/fasthttp"
)

// PreFlag indicates whether Handlers.Handler and Handlers.PostHandler are to be skipped.
type PreFlag int

// Various flag for Handlers.PreHandler
const (
	Continue PreFlag = iota
	Stop
)

// RequestHandler is fasthttp.RequestHandler with error return.
type RequestHandler func(ctx *fasthttp.RequestCtx) PreFlag

// Handlers is a collection of request handlers.
type Handlers struct {
	PreHandler   []RequestHandler
	Handler      fasthttp.RequestHandler
	PostHandler  []fasthttp.RequestHandler
	FinalHandler []fasthttp.RequestHandler
}

// AddPreHandler add a RequestHandler to the end of PreHandler slice.
func (h *Handlers) AddPreHandler(r RequestHandler) {
	if h.PreHandler == nil {
		h.PreHandler = []RequestHandler{}
	}
	h.PreHandler = append(h.PreHandler, r)
}

// AddPostHandler add a RequestHandler to the end of PostHandler slice.
func (h *Handlers) AddPostHandler(r fasthttp.RequestHandler) {
	if h.PostHandler == nil {
		h.PostHandler = []fasthttp.RequestHandler{}
	}
	h.PostHandler = append(h.PostHandler, r)
}

// AddFinalHandler add a RequestHandler to the end of FinalHandler slice.
func (h *Handlers) AddFinalHandler(r fasthttp.RequestHandler) {
	if h.FinalHandler == nil {
		h.FinalHandler = []fasthttp.RequestHandler{}
	}
	h.FinalHandler = append(h.FinalHandler, r)
}

// CopyHandlers copy Handlers, such that you can add more specific handlers.
func (h *Handlers) CopyHandlers() *Handlers {
	c := &Handlers{}
	if h.PreHandler != nil {
		c.PreHandler = make([]RequestHandler, len(h.PreHandler))
		copy(c.PreHandler, h.PreHandler)
	}
	if h.PostHandler != nil {
		c.PostHandler = make([]fasthttp.RequestHandler, len(h.PostHandler))
		copy(c.PostHandler, h.PostHandler)
	}
	if h.FinalHandler != nil {
		c.FinalHandler = make([]fasthttp.RequestHandler, len(h.FinalHandler))
		copy(c.FinalHandler, h.FinalHandler)
	}
	return c
}

// SetHandler return a copy of Handlers with its Handler are set.
func (h *Handlers) SetHandler(r fasthttp.RequestHandler) *Handlers {
	c := h.CopyHandlers()
	c.Handler = r
	return c
}
