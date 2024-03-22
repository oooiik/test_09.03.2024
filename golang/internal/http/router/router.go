package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	mux *gin.Engine
}

func New() *Router {
	r := Router{
		mux: gin.Default(),
	}
	return &r
}

func (r *Router) Handler() http.Handler {
	return r.mux
}
