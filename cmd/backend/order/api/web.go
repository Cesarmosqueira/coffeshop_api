package api

import (
	"net/http"

	o "github.com/Cesarmosqueira/coffeshop_api/cmd/backend/order"
	r "github.com/Cesarmosqueira/coffeshop_api/pkg/domain/response"
	"github.com/gin-gonic/gin"
)

type webApi struct {
	router  *gin.RouterGroup
	service o.OrderService
}

type OrderWebApi interface {
	CreateOrder()
	GetById()
	ListOrders()
	DeleteById()
}

func NewOrderWebApi(router *gin.Engine) OrderWebApi {
	return &webApi{
		router:  router.Group("/api/orders"),
		service: o.NewOrderService(),
	}
}

func (a *webApi) CreateOrder() {
	a.router.POST("", func(c *gin.Context) {
		var body o.OrderDto
		if err := c.ShouldBindJSON(&body); err != nil {
			c.AbortWithStatusJSON(r.NewResponse(http.StatusInternalServerError, "Cannot unmarshal the object"))
			return
		}

		validationErrors := body.Validate()
		if len(validationErrors) > 0 {
			c.AbortWithStatusJSON(r.NewValidationErrorResponse(validationErrors))
			return
		}

		order, err := a.service.CreateOrder(body)


		if err != nil {
			c.AbortWithStatusJSON(r.NewResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		c.JSON(http.StatusOK, order)
	})
}

func (a *webApi) ListOrders() {
	a.router.GET("", func(c *gin.Context) {
		orders, err := a.service.ListOrders()

		if err != nil {
			c.AbortWithStatusJSON(r.NewResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		c.JSON(http.StatusOK, orders)
	})
}

func (a *webApi) GetById() {
	a.router.GET(":orderid", func(c *gin.Context) {
		orderid := c.Param("orderid")


		order, err := a.service.GetOrder(orderid)

		if err != nil {
			c.AbortWithStatusJSON(r.NewResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		c.JSON(http.StatusOK, order)
	})
}

func (a *webApi) DeleteById() {
	a.router.DELETE(":orderid", func(c *gin.Context) {
		orderid := c.Param("orderid")


		count, err := a.service.DeleteOrder(orderid)

		if err != nil {
			c.AbortWithStatusJSON(r.NewResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		c.JSON(http.StatusOK, count)
	})
}
