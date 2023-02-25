package handler

import (
	"codingTest/app/helper"
	reviewproduct "codingTest/app/reviewProduct"
	"net/http"

	"github.com/gin-gonic/gin"
)

type reviewProductHandler struct {
	service reviewproduct.Service
}

func NewReviewProductHandler(revProductService reviewproduct.Service) *reviewProductHandler {
	return &reviewProductHandler{revProductService}
}

func (h *reviewProductHandler) GetAllReviewProducts(ctx *gin.Context) {
	allRevProducts, err := h.service.GetReviewProduct()

	if err != nil {
		response := helper.ApiResponse("error to get review products", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("list of products", http.StatusOK, "success", reviewproduct.FormatRProductSlice(allRevProducts))

	ctx.JSON(http.StatusOK, response)
}
