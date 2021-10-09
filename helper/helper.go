package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	hasil := Response{
		Meta: meta,
		Data: data,
	}

	return hasil

}

func FormatValidationError(err error) []string {
	var errors []string
	tempError := ""
	for _, v := range err.(validator.ValidationErrors) {
		switch v.Tag() {
		case "required":
			tempError = fmt.Sprintf("%s is required", v.Field())
		case "email":
			tempError = fmt.Sprintf("%s is not valid email", v.Field())
		case "gte":
			tempError = fmt.Sprintf("%s value must be greater than %s", v.Field(), v.Param())
		case "lte":
			tempError = fmt.Sprintf("%s value must be lower than %s", v.Field(), v.Param())
		case "min":
			tempError = fmt.Sprintf("%s character must be min %s", v.Field(), v.Param())
		case "max":
			tempError = fmt.Sprintf("%s character must be max %s", v.Field(), v.Param())
		default:
			tempError = v.Error()

		}
		errors = append(errors, tempError)
	}
	return errors
}
