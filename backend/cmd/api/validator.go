// Package api (validator) defines custom validations to improve the error message.
package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// validationError represents a custom validation error that contains information about the violated fields and their messages.
type validationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// bindJSONWithValidation is a helper function that binds the JSON request body to the given interface and validates it with the specified validator.
func bindJSONWithValidation(ctx *gin.Context, req interface{}, validate *validator.Validate) error {
	if err := ctx.ShouldBindJSON(req); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}

		validationErrors := make([]validationError, 0, len(errs))

		for _, err := range errs {
			validationErrors = append(validationErrors, validationError{
				Field: err.Field(),
				Error: fmt.Sprintf("%s validation failed on '%s'", err.Tag(), err.Param()),
			})
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": validationErrors,
		})
		return err
	}

	return validate.Struct(req)
}
