package helper

import (
	"codingTest/config"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) interface{} {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	errorMessage := gin.H{"errors": errors}
	return errorMessage
}

func ErrorValidation(err error, c *gin.Context, msg string, status string, code int, errMsg interface{}) {
	byteData := InterfaceToByte(errMsg)
	log := ApiResponse(msg, http.StatusUnprocessableEntity, status, string(byteData))
	config.Loggers("error", log)

	response := ApiResponse(msg, http.StatusUnprocessableEntity, status, errMsg)
	c.JSON(code, response)
}

func MapToByte(params map[string]string) []byte {
	jsonContent, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		return jsonContent
	}
	return jsonContent
}

func InterfaceToByte(params interface{}) []byte {
	jsonContent, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		return jsonContent
	}
	return jsonContent
}
