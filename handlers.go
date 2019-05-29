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
	Checkpoint   []CheckHandler
	Handler      fasthttp.RequestHandler
	PostHandler  []fasthttp.RequestHandler
	FinalHandler []fasthttp.RequestHandler
}

// AddCheckpoint add a CheckHandler to the end of Checkpoint slice.
func (h *Handlers) AddCheckpoint(r CheckHandler) {
	if h.Checkpoint == nil {
		h.Checkpoint = []CheckHandler{}
	}
	h.Checkpoint = append(h.Checkpoint, r)
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
	if h.Checkpoint != nil {
		c.Checkpoint = make([]CheckHandler, len(h.Checkpoint))
		copy(c.Checkpoint, h.Checkpoint)
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
