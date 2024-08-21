package handlers

import (
	"net/http"

	"github.com/adafatya/micro-services/webapi/internal/dto"
	"github.com/adafatya/micro-services/webapi/internal/services"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: productService,
	}
}

func (p *ProductHandler) AddProduct(c *gin.Context) {
	var req dto.AddProductRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.MessageResponse{Message: "Gagal menambah produk: data tidak sesuai!"})
		return
	}

	resp, err := p.ProductService.AddProduct(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
	}

	c.JSON(http.StatusOK, resp)
}
