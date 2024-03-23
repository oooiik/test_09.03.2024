package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/oooiik/test_09.03.2024/internal/http/request"
	"github.com/oooiik/test_09.03.2024/internal/http/response"
	"github.com/oooiik/test_09.03.2024/internal/logger"
	"github.com/oooiik/test_09.03.2024/internal/service"
	"net/http"
)

type Good interface {
	Index(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	RePrioritize(ctx *gin.Context)
}

type good struct {
	service service.Good
}

func NewGood(service service.Good) Good {
	return &good{
		service: service,
	}
}

func (c good) Index(ctx *gin.Context) {
	req := request.GoodIndex{Limit: 10, Offset: 0}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.Err(err))
		return
	}

	list, err := c.service.Index(req)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, response.Err(err))
		return
	}

	meta, err := c.service.Meta()
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, response.Err(err))
		return
	}
	meta.Offset = req.Offset
	meta.Limit = req.Limit

	ctx.JSON(http.StatusOK, response.GoodListResponse(list, meta))
}

func (c good) Create(ctx *gin.Context) {
	var req request.GoodCreate
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.Err(err))
		return
	}

	model, err := c.service.Create(req)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, response.Err(err))
		return
	}

	ctx.JSON(http.StatusOK, response.GoodResponse(model))
}

func (c good) Update(ctx *gin.Context) {
	var req request.GoodUpdate

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.Err(err))
		return
	}

	model, err := c.service.Update(req)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, response.Err(err))
		return
	}

	ctx.JSON(http.StatusOK, response.GoodResponse(model))
}

func (c good) Delete(ctx *gin.Context) {
	var req request.GoodDelete
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.Err(err))
		return
	}

	model, err := c.service.Delete(req.Id)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, response.Err(err))
		return
	}

	ctx.JSON(http.StatusOK, response.GoodDeletedResponse(model))
}

func (c good) RePrioritize(ctx *gin.Context) {
	var req request.GoodRePrioritize

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.Err(err))
		return
	}

	list, err := c.service.RePrioritize(req)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, response.Err(err))
		return
	}

	ctx.JSON(http.StatusOK, response.GoodListRePrioritiesResponse(list))
}
