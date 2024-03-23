package router

import "github.com/oooiik/test_09.03.2024/internal/http/controller"

func (r *Router) ApiGoods(controller controller.Good) {
	r.mux.GET("/goods", controller.Index)
	r.mux.POST("/goods/create", controller.Create)
	r.mux.PATCH("/goods/update", controller.Update)
}
