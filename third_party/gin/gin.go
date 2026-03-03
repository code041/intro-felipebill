package gin

import (
	"encoding/json"
	"net/http"
)

type H map[string]any

type HandlerFunc func(*Context)

type route struct {
	method  string
	path    string
	handler HandlerFunc
}

type Engine struct {
	routes []route
}

func Default() *Engine {
	return &Engine{}
}

func (e *Engine) GET(path string, handler HandlerFunc) {
	e.routes = append(e.routes, route{method: http.MethodGet, path: path, handler: handler})
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, rt := range e.routes {
		if rt.method == r.Method && rt.path == r.URL.Path {
			rt.handler(&Context{Request: r, Writer: w})
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter
}

func (c *Context) JSON(code int, obj any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(code)
	_ = json.NewEncoder(c.Writer).Encode(obj)
}

func (c *Context) DefaultQuery(key, defaultValue string) string {
	value := c.Request.URL.Query().Get(key)
	if value == "" {
		return defaultValue
	}
	return value
}
