package response

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

const (
	ErrCodeNoSuchUser        = 1
	ErrTooManyFailedAttempts = 2
	ErrCodeLoginBlocked      = 3
)

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type CustomError struct {
	Code             int               `json:"code"`
	Message          string            `json:"message"`
	DevMessage       string            `json:"dev_message"`
	ValidationErrors []ValidationError `json:"validation_errors,omitempty"`
}

func InternalError(c *gin.Context, err error) {
	ErrorResponse(c, err, http.StatusInternalServerError, "internal error")
}

func BadRequestPath(c *gin.Context, err error) {
	ErrorResponse(c, err, http.StatusBadRequest, "bad request path")
}

func BadRequestBody(c *gin.Context, err error) {
	ErrorResponse(c, err, http.StatusBadRequest, "bad request body")
}

func ErrorResponse(c *gin.Context, err error, code int, msg string) {
	if err == nil {
		err = errors.New("placeholder error")
	}

	customErr := CustomError{
		Code:       code,
		Message:    msg,
		DevMessage: err.Error(),
	}

	var fieldErrs validator.ValidationErrors
	ok := errors.As(err, &fieldErrs)
	if ok {
		for _, fieldErr := range fieldErrs {
			customErr.ValidationErrors = append(customErr.ValidationErrors, ValidationError{
				Field: toSnakeCase(fieldErr.Field()),
				Error: errFromTag(fieldErr.Tag()),
			})
		}
	}

	c.JSON(code, customErr)
}

func toSnakeCase(camel string) string {
	var buf bytes.Buffer
	for _, c := range camel {
		if 'A' <= c && c <= 'Z' {
			if buf.Len() > 0 {
				buf.WriteRune('_')
			}
			buf.WriteRune(c - 'A' + 'a')
		} else {
			buf.WriteRune(c)
		}
	}
	return buf.String()
}

func errFromTag(tag string) string {
	switch tag {
	case "required":
		return "field is required"
	default:
		return fmt.Sprintf("invalid %s", tag)
	}
}
