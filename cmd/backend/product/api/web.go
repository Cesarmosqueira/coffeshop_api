package api

import (
	"net/http"

	p "github.com/Cesarmosqueira/coffeshop_api/cmd/backend/product"
	r "github.com/Cesarmosqueira/coffeshop_api/pkg/domain/response"
	"github.com/gin-gonic/gin"
)

type webApi struct {
	router  *gin.RouterGroup
	service p.ProductService
}

type ProductWebApi interface {
	CreateProduct()
	GetById()
	GetByProductCode()
	ListProducts()
	DeleteById()
}

func NewProductWebApi(router *gin.Engine) ProductWebApi {
	return &webApi{
		router:  router.Group("/api/products"),
		service: p.NewProductService(),
	}
}

func (a *webApi) CreateProduct() {
	a.router.POST("", func(c *gin.Context) {
		var body p.ProductDto
		if err := c.ShouldBindJSON(&body); err != nil {
			c.AbortWithStatusJSON(r.NewResponse(http.StatusInternalServerError, "Cannot unmarshal the object"))
			return
		}

		validationErrors := body.Validate()
		if len(validationErrors) > 0 {
			c.AbortWithStatusJSON(r.NewValidationErrorResponse(validationErrors))
			return
		}

		product, err := a.service.CreateProduct(body)


		if err != nil {
			c.AbortWithStatusJSON(r.NewResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		c.JSON(http.StatusOK, product)
	})
}

func (a *webApi) ListProducts() {
	a.router.GET("", func(c *gin.Context) {
		products, err := a.service.ListProducts()

		if err != nil {
			c.AbortWithStatusJSON(r.NewResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		c.JSON(http.StatusOK, products)
	})
}

func (a *webApi) GetById() {
	a.router.GET(":productid", func(c *gin.Context) {
		productid := c.Param("productid")


		product, err := a.service.GetProduct(productid)

		if err != nil {
			c.AbortWithStatusJSON(r.NewResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		c.JSON(http.StatusOK, product)
	})
}

func (a *webApi) GetByProductCode() {
	a.router.GET("code/:productcode", func(c *gin.Context) {
		productCode := c.Param("productcode")


		product, err := a.service.GetByProductCode(productCode)

		if err != nil {
			c.AbortWithStatusJSON(r.NewResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		c.JSON(http.StatusOK, product)
	})
}

func (a *webApi) DeleteById() {
	a.router.DELETE(":productid", func(c *gin.Context) {
		productid := c.Param("productid")


		count, err := a.service.DeleteProduct(productid)

		if err != nil {
			c.AbortWithStatusJSON(r.NewResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		c.JSON(http.StatusOK, count)
	})
}
