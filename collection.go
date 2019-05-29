package router

import (
	"strings"

	"github.com/valyala/fasthttp"
)

// Collection contain a Router and a Handlers.
type Collection struct {
	Router   *Router
	Handlers *Handlers
}

// NewCollection return a new empty Collection with router set and empty Handlers.
func NewCollection() *Collection {
	return &Collection{
		Router:   New(),
		Handlers: &Handlers{},
	}
}

// Copy return a new Collection with Router point to the same Router but handlers are copied.
func (c *Collection) Copy() *Collection {
	n := &Collection{}
	n.Router = c.Router
	n.Handlers = c.Handlers.CopyHandlers()
	return n
}

// GET is a shortcut for Collection.Handle("GET", path, handle).
func (c *Collection) GET(path string, handle fasthttp.RequestHandler) {
	c.Router.Handle("GET", path, c.Handlers.SetHandler(handle))
}

// HEAD is a shortcut for Collection.Handle("HEAD", path, handle).
func (c *Collection) HEAD(path string, handle fasthttp.RequestHandler) {
	c.Router.Handle("HEAD", path, c.Handlers.SetHandler(handle))
}

// OPTIONS is a shortcut for Collection.Handle("OPTIONS", path, handle).
func (c *Collection) OPTIONS(path string, handle fasthttp.RequestHandler) {
	c.Router.Handle("OPTIONS", path, c.Handlers.SetHandler(handle))
}

// POST is a shortcut for Collection.Handle("POST", path, handle).
func (c *Collection) POST(path string, handle fasthttp.RequestHandler) {
	c.Router.Handle("POST", path, c.Handlers.SetHandler(handle))
}

// PUT is a shortcut for Collection.Handle("PUT", path, handle).
func (c *Collection) PUT(path string, handle fasthttp.RequestHandler) {
	c.Router.Handle("PUT", path, c.Handlers.SetHandler(handle))
}

// PATCH is a shortcut for Collection.Handle("PATCH", path, handle).
func (c *Collection) PATCH(path string, handle fasthttp.RequestHandler) {
	c.Router.Handle("PATCH", path, c.Handlers.SetHandler(handle))
}

// DELETE is a shortcut for Collection.Handle("DELETE", path, handle).
func (c *Collection) DELETE(path string, handle fasthttp.RequestHandler) {
	c.Router.Handle("DELETE", path, c.Handlers.SetHandler(handle))
}

// Handle registers a new request handlers with the given path and method.
//
// For GET, POST, PUT, PATCH and DELETE requests the respective shortcut
// functions can be used.
func (c *Collection) Handle(method, path string, handle fasthttp.RequestHandler) {
	c.Router.Handle(method, path, c.Handlers.SetHandler(handle))
}

// ServeFiles serves files from the given file system root.
// The path must end with "/*filepath", files are then served from the local
// path /defined/root/dir/*filepath.
// For example if root is "/etc" and *filepath is "passwd", the local file
// "/etc/passwd" would be served.
// Internally a http.FileServer is used, therefore http.NotFound is used instead
// of the Router's NotFound handler.
//     router.ServeFiles("/src/*filepath", "/var/www")
func (c *Collection) ServeFiles(path string, rootPath string) {
	if len(path) < 10 || path[len(path)-10:] != "/*filepath" {
		panic("path must end with /*filepath in path '" + path + "'")
	}
	prefix := path[:len(path)-10]

	fileHandler := fasthttp.FSHandler(rootPath, strings.Count(prefix, "/"))

	c.GET(path, func(ctx *fasthttp.RequestCtx) {
		fileHandler(ctx)
	})
}
