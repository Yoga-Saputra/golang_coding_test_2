package handler

import (
	"codingTest/app/helper"
	"codingTest/app/members"
	reviewproduct "codingTest/app/reviewProduct"
	"net/http"

	"github.com/gin-gonic/gin"
)

type memberHandler struct {
	memberService     members.Service
	revProductService reviewproduct.Service
}

func NewMemberHandler(memberService members.Service, revProductService reviewproduct.Service) *memberHandler {
	return &memberHandler{memberService, revProductService}
}

func (h *memberHandler) GetALlMembers(ctx *gin.Context) {
	allMembers, err := h.memberService.GetMembers()

	if err != nil {
		response := helper.ApiResponse("error to get members", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("list of members", http.StatusOK, "success", members.FormatMemberSlice(allMembers))

	ctx.JSON(http.StatusOK, response)
}

func (h *memberHandler) GetMember(ctx *gin.Context) {
	var input members.GetMemberDetailInput

	err := ctx.ShouldBindUri(&input)

	if err != nil {
		response := helper.ApiResponse("failed to get detail of member", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	member, err := h.memberService.GetMemberById(input.IDMember)

	if err != nil {
		response := helper.ApiResponse("failed to get detail of member", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := members.FormatMember(member)
	response := helper.ApiResponse("member detail", http.StatusOK, "success", formatter)
	ctx.JSON(http.StatusOK, response)
}

func (h *memberHandler) Save(ctx *gin.Context) {
	var input members.InputMember
	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		helper.ErrorValidation(err, ctx, "insert member failed", "error", http.StatusUnprocessableEntity, errors)
		return
	}

	membersRes, err := h.memberService.InsertMember(input)

	if err != nil {
		helper.ErrorValidation(err, ctx, "insert member failed", "error", http.StatusUnprocessableEntity, err.Error())
		return
	}

	formatter := members.FormatMember(membersRes)
	response := helper.ApiResponse("member has been created", http.StatusOK, "success", formatter)

	ctx.JSON(http.StatusOK, response)
}

func (h *memberHandler) Update(ctx *gin.Context) {
	var inputId members.GetMemberDetailInput

	err := ctx.ShouldBindUri(&inputId)

	if err != nil {
		response := helper.ApiResponse("failed to update member", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData members.InputMember

	err = ctx.ShouldBindJSON(&inputData)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		helper.ErrorValidation(err, ctx, "failed to update member", "error", http.StatusBadRequest, errorMessage)
		return
	}

	updateMember, err := h.memberService.UpdateMember(inputData, inputId.IDMember)

	if err != nil {
		response := helper.ApiResponse("failed to update member", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := members.FormatMember(updateMember)
	response := helper.ApiResponse("success to update member", http.StatusOK, "success", formatter)
	ctx.JSON(http.StatusOK, response)

}

func (h *memberHandler) Delete(ctx *gin.Context) {
	var input members.GetMemberDetailInput

	err := ctx.ShouldBindUri(&input)

	if err != nil {
		response := helper.ApiResponse("failed to delete member", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	rProduct, err := h.revProductService.GetReviewProductById(input.IDMember, "id_member")
	if err != nil {
		response := helper.ApiResponse("failed to delete member", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if rProduct.ID_Review > 0 {
		msg := "Cannot delete a parent row: a foreign key constraint that id_member"
		response := helper.ApiResponse("failed to delete member", http.StatusForbidden, "error", msg)
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	_, err = h.memberService.DeleteMember(input.IDMember)

	if err != nil {
		response := helper.ApiResponse("failed to delete member", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("member detail", http.StatusOK, "success", "member already deleted")
	ctx.JSON(http.StatusOK, response)
}
