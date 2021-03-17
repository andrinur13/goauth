package handler

import (
	"goauth/helper"
	"goauth/product"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *productHandler {
	return &productHandler{productService}
}

func (h *productHandler) InsertProduct(c *gin.Context) {
	var input product.ProductInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		formatter := helper.APIResponse("failed", http.StatusUnprocessableEntity, "failed to insert product", err.Error())
		c.JSON(http.StatusUnprocessableEntity, formatter)
		return
	}

	// save ke db
	data, err := h.productService.SaveProduct(input)

	if err != nil {
		formatter := helper.APIResponse("failed", http.StatusUnprocessableEntity, "failed to insert product", err.Error())
		c.JSON(http.StatusUnprocessableEntity, formatter)
		return
	}

	// mapping ke output product
	output := helper.OutputProduct{
		Name:  data.Name,
		Price: data.Price,
	}

	// return jika berhasil
	formatter := helper.APIResponse("success", http.StatusOK, "success insert product", output)
	c.JSON(http.StatusOK, formatter)
}

func (h *productHandler) GetAllProducts(c *gin.Context) {
	// ambil data dari service
	data, err := h.productService.GetProducts()

	if err != nil {
		formatter := helper.APIResponse("failed", http.StatusUnprocessableEntity, "failed to get Products", err.Error())
		c.JSON(http.StatusUnprocessableEntity, formatter)
		return
	}

	productFormatter := product.FormatOutputProducts(data)
	formatter := helper.APIResponse("success", http.StatusOK, "success get all products", productFormatter)
	c.JSON(http.StatusOK, formatter)
}
