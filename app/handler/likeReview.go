package handler

import (
	"codingTest/app/helper"
	likereview "codingTest/app/likeReview"
	"codingTest/app/members"
	reviewproduct "codingTest/app/reviewProduct"
	"net/http"

	"github.com/gin-gonic/gin"
)

type likeReviewHandler struct {
	lreviewService    likereview.Service
	revProductService reviewproduct.Service
	memberService     members.Service
}

func NewLikeReviewHandler(lreviewService likereview.Service, revProductService reviewproduct.Service, memberService members.Service) *likeReviewHandler {
	return &likeReviewHandler{lreviewService, revProductService, memberService}
}

func (h *likeReviewHandler) LikeRev(ctx *gin.Context) {
	var input likereview.LikeReview
	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		helper.ErrorValidation(err, ctx, "like review failed", "error", http.StatusUnprocessableEntity, errors)
		return
	}

	// check member exists
	_, err = h.memberService.GetMemberById(input.ID_Member)

	if err != nil {
		helper.ErrorValidation(err, ctx, "insert review product failed", "error", http.StatusUnprocessableEntity, err.Error())
		return
	}

	// check review product exists
	_, err = h.revProductService.GetReviewProductById(input.ID_Review, "id_review")
	if err != nil {
		helper.ErrorValidation(err, ctx, "insert review product failed", "error", http.StatusUnprocessableEntity, err.Error())
		return
	}

	membersRes, err := h.lreviewService.Like(input)

	if err != nil {
		helper.ErrorValidation(err, ctx, "insert member failed", "error", http.StatusUnprocessableEntity, err.Error())
		return
	}

	response := helper.ApiResponse("member has been created", http.StatusOK, "success", membersRes)

	ctx.JSON(http.StatusOK, response)
}

func (h *likeReviewHandler) Dislike(ctx *gin.Context) {
	var input likereview.LikeReview
	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		helper.ErrorValidation(err, ctx, "dislike review failed", "error", http.StatusUnprocessableEntity, errors)
		return
	}

	_, err = h.lreviewService.DisLike(input.ID_Review, input.ID_Member)
	if err != nil {
		response := helper.ApiResponse("failed to dislike review", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("dislike detail", http.StatusOK, "success", "already dislike review")
	ctx.JSON(http.StatusOK, response)
}
