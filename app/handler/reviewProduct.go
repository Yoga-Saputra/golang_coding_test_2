package handler

import (
	"codingTest/app/helper"
	"codingTest/app/members"
	"codingTest/app/products"
	reviewproduct "codingTest/app/reviewProduct"
	"net/http"

	"github.com/gin-gonic/gin"
)

type reviewProductHandler struct {
	revProductService reviewproduct.Service
	productService    products.Service
	memberService     members.Service
}

func NewReviewProductHandler(revProductService reviewproduct.Service, productService products.Service, memberService members.Service) *reviewProductHandler {
	return &reviewProductHandler{revProductService, productService, memberService}
}

func (h *reviewProductHandler) GetAllReviewProducts(ctx *gin.Context) {
	allRevProducts, err := h.revProductService.GetReviewProduct()

	if err != nil {
		response := helper.ApiResponse("error to get review products", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("list of products", http.StatusOK, "success", reviewproduct.FormatRProductSlice(allRevProducts))

	ctx.JSON(http.StatusOK, response)
}

func (h *reviewProductHandler) Save(ctx *gin.Context) {
	var input reviewproduct.InputReviewProduct
	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		helper.ErrorValidation(err, ctx, "insert review product failed", "error", http.StatusUnprocessableEntity, errors)
		return
	}
	// check member exists
	_, err = h.memberService.GetMemberById(input.ID_Memebr)

	if err != nil {
		helper.ErrorValidation(err, ctx, "insert review product failed", "error", http.StatusUnprocessableEntity, err.Error())
		return
	}

	// check product exists
	_, err = h.productService.GetProductById(input.ID_Product)
	if err != nil {
		helper.ErrorValidation(err, ctx, "insert review product failed", "error", http.StatusUnprocessableEntity, err.Error())
		return
	}

	rProductsRes, err := h.revProductService.InsertReviewProduct(input)

	if err != nil {
		helper.ErrorValidation(err, ctx, "insert review product failed", "error", http.StatusUnprocessableEntity, err.Error())
		return
	}

	formatter := reviewproduct.FormatRProductCreate(rProductsRes)
	response := helper.ApiResponse("review product has been created", http.StatusOK, "success", formatter)

	ctx.JSON(http.StatusOK, response)
}
