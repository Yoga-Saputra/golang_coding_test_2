package handler

import (
	"codingTest/app/helper"
	"codingTest/app/members"
	"net/http"

	"github.com/gin-gonic/gin"
)

type memberHandler struct {
	memberService members.Service
}

func NewMemberHandler(memberService members.Service) *memberHandler {
	return &memberHandler{memberService}
}

func (h *memberHandler) GetALlMembers(ctx *gin.Context) {
	allMembers, err := h.memberService.GetMembers()

	if err != nil {
		response := helper.ApiResponse("error to get campaigns", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("list of campaigns", http.StatusOK, "success", members.FormatMemberSlice(allMembers))

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

	member, err := h.memberService.GetMemberById(input)

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

	updateMember, err := h.memberService.UpdateMember(inputData, inputId)

	if err != nil {
		response := helper.ApiResponse("failed to update member", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := members.FormatMember(updateMember)
	response := helper.ApiResponse("success to update campaign", http.StatusOK, "success", formatter)
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

	_, err = h.memberService.DeleteMember(input)

	if err != nil {
		response := helper.ApiResponse("failed to delete member", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("member detail", http.StatusOK, "success", "member already deleted")
	ctx.JSON(http.StatusOK, response)
}
