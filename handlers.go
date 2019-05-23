package router

import (
	"github.com/valyala/fasthttp"
)

// RequestHandler is fasthttp.RequestHandler with error return.
type RequestHandler func(ctx *fasthttp.RequestCtx) error

// Handlers is a collection of request handlers.
type Handlers struct {
	preHandler   []RequestHandler
	handler      fasthttp.RequestHandler
	postHandler  []fasthttp.RequestHandler
	finalHandler []fasthttp.RequestHandler
}

// AddPreHandler add a RequestHandler to the end of preHandler slice.
func (h *Handlers) AddPreHandler(r RequestHandler) {
	if h.preHandler == nil {
		h.preHandler = []RequestHandler{}
	}
	h.preHandler = append(h.preHandler, r)
}

// AddPostHandler add a RequestHandler to the end of postHandler slice.
func (h *Handlers) AddPostHandler(r fasthttp.RequestHandler) {
	if h.postHandler == nil {
		h.postHandler = []fasthttp.RequestHandler{}
	}
	h.postHandler = append(h.postHandler, r)
}

// AddFinalHandler add a RequestHandler to the end of postHandler slice.
func (h *Handlers) AddFinalHandler(r fasthttp.RequestHandler) {
	if h.finalHandler == nil {
		h.finalHandler = []fasthttp.RequestHandler{}
	}
	h.finalHandler = append(h.finalHandler, r)
}

// CopyHandlers copy Handlers, such that you can add more specific handlers.
func (h *Handlers) CopyHandlers() *Handlers {
	c := &Handlers{}
	if h.preHandler != nil {
		c.preHandler = make([]RequestHandler, len(h.preHandler))
		copy(c.preHandler, h.preHandler)
	}
	if h.postHandler != nil {
		c.postHandler = make([]fasthttp.RequestHandler, len(h.postHandler))
		copy(c.postHandler, h.postHandler)
	}
	if h.finalHandler != nil {
		c.finalHandler = make([]fasthttp.RequestHandler, len(h.finalHandler))
		copy(c.finalHandler, h.finalHandler)
	}
	return c
}

// SetHandler return a copy of Handlers with its handler are set.
func (h *Handlers) SetHandler(r fasthttp.RequestHandler) *Handlers {
	c := h.CopyHandlers()
	c.handler = r
	return c
}
