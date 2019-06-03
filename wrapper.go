package router

import (
	"github.com/valyala/fasthttp"
)

// Wrapper is a functions wrapper.
type Wrapper struct {
	Checkpoint   []CheckHandler
	PostHandler  []fasthttp.RequestHandler
	FinalHandler []fasthttp.RequestHandler
}

// AddCheckpoint add a CheckHandler to the Wrapper.
func (w *Wrapper) AddCheckpoint(r CheckHandler) {
	w.Checkpoint = append(w.Checkpoint, r)
}

// AddPostHandler add a RequestHandler to the end of PostHandler slice.
func (w *Wrapper) AddPostHandler(r fasthttp.RequestHandler) {
	w.PostHandler = append(w.PostHandler, r)
}

// AddFinalHandler add a RequestHandler to the end of FinalHandler slice.
func (w *Wrapper) AddFinalHandler(r fasthttp.RequestHandler) {
	w.FinalHandler = append(w.FinalHandler, r)
}

// Copy copies all functions wrapper.
func (w *Wrapper) Copy() *Wrapper {
	c := &Wrapper{}
	if w.Checkpoint != nil {
		c.Checkpoint = make([]CheckHandler, len(w.Checkpoint))
		copy(c.Checkpoint, w.Checkpoint)
	}
	if w.PostHandler != nil {
		c.PostHandler = make([]fasthttp.RequestHandler, len(w.PostHandler))
		copy(c.PostHandler, w.PostHandler)
	}
	if w.FinalHandler != nil {
		c.FinalHandler = make([]fasthttp.RequestHandler, len(w.FinalHandler))
		copy(c.FinalHandler, w.FinalHandler)
	}
	return c
}
