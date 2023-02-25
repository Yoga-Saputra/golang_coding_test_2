package handler

import (
	"codingTest/app/helper"
	"codingTest/app/products"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService products.Service
}

func NewProductHandler(productService products.Service) *productHandler {
	return &productHandler{productService}
}

func (h *productHandler) GetALlProduct(ctx *gin.Context) {
	allProducts, err := h.productService.GetProducts()

	if err != nil {
		response := helper.ApiResponse("error to get products", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("list of products", http.StatusOK, "success", products.FormatProductSlice(allProducts))

	ctx.JSON(http.StatusOK, response)
}

func (h *productHandler) GetProduct(ctx *gin.Context) {
	var input products.GetProductDetailInput

	err := ctx.ShouldBindUri(&input)

	if err != nil {
		response := helper.ApiResponse("failed to get detail of product", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := h.productService.GetProductById(input)

	if err != nil {
		response := helper.ApiResponse("failed to get detail of product", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := products.FormatProduct(product)
	response := helper.ApiResponse("product detail", http.StatusOK, "success", formatter)
	ctx.JSON(http.StatusOK, response)
}

func (h *productHandler) Save(ctx *gin.Context) {
	var input products.InputProduct
	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		helper.ErrorValidation(err, ctx, "insert product failed", "error", http.StatusUnprocessableEntity, errors)
		return
	}

	productsRes, err := h.productService.InsertProduct(input)

	if err != nil {
		helper.ErrorValidation(err, ctx, "insert product failed", "error", http.StatusUnprocessableEntity, err.Error())
		return
	}

	formatter := products.FormatProduct(productsRes)
	response := helper.ApiResponse("product has been created", http.StatusOK, "success", formatter)

	ctx.JSON(http.StatusOK, response)
}

func (h *productHandler) Update(ctx *gin.Context) {
	var inputId products.GetProductDetailInput

	err := ctx.ShouldBindUri(&inputId)

	if err != nil {
		response := helper.ApiResponse("failed to update product", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData products.InputProduct

	err = ctx.ShouldBindJSON(&inputData)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		helper.ErrorValidation(err, ctx, "failed to update product", "error", http.StatusBadRequest, errorMessage)
		return
	}

	updateproduct, err := h.productService.UpdateProduct(inputData, inputId)

	if err != nil {
		response := helper.ApiResponse("failed to update product", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := products.FormatProduct(updateproduct)
	response := helper.ApiResponse("success to update product", http.StatusOK, "success", formatter)
	ctx.JSON(http.StatusOK, response)

}

func (h *productHandler) Delete(ctx *gin.Context) {
	var input products.GetProductDetailInput

	err := ctx.ShouldBindUri(&input)

	if err != nil {
		response := helper.ApiResponse("failed to delete product", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.productService.DeleteProduct(input)

	if err != nil {
		response := helper.ApiResponse("failed to delete product", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("product detail", http.StatusOK, "success", "product already deleted")
	ctx.JSON(http.StatusOK, response)
}
