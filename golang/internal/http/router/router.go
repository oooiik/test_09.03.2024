package router

import (
	"net/http"
)

type Router struct {
	mux *http.ServeMux
}

func New() *Router {
	return &Router{
		mux: &http.ServeMux{},
	}
}

func (r *Router) Handler() *http.ServeMux {
	return r.mux
}
