package router

import "net/http"

type Group struct {
	mux        *http.ServeMux
	prefix     string
	middleware func(http.Handler) http.Handler
}

func NewGroup(mux *http.ServeMux, prefix string, middleware func(http.Handler) http.Handler) *Group {
	return &Group{
		mux:        mux,
		prefix:     prefix,
		middleware: middleware,
	}
}

func (g *Group) Handle(method, path string, handlerFunc http.HandlerFunc) {
	h := http.Handler(handlerFunc)

	if g.middleware != nil {
		h = g.middleware(h)
	}

	fullPath := g.prefix + path

	g.mux.HandleFunc(fullPath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func (g *Group) Get(path string, handlerFunc http.HandlerFunc) {
	g.Handle(http.MethodGet, path, handlerFunc)
}

func (g *Group) Post(path string, handlerFunc http.HandlerFunc) {
	g.Handle(http.MethodPost, path, handlerFunc)
}

func (g *Group) Put(path string, handlerFunc http.HandlerFunc) {
	g.Handle(http.MethodPut, path, handlerFunc)
}

func (g *Group) Patch(path string, handlerFunc http.HandlerFunc) {
	g.Handle(http.MethodPatch, path, handlerFunc)
}

func (g *Group) Delete(path string, handlerFunc http.HandlerFunc) {
	g.Handle(http.MethodDelete, path, handlerFunc)
}
