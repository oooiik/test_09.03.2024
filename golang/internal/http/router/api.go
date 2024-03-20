package router

import "github.com/oooiik/test_09.03.2024/internal/http/controller"

func (r *Router) api() {
	r.mux.GET("/goods", controller.SingletonGood().Index)
}
