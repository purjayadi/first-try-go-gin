package middleware

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateInputMiddleware(inputStruct interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut || c.Request.Method == http.MethodPatch {
			if err := c.ShouldBindJSON(inputStruct); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
		}

		// Check if the request contains query parameters
		if c.Request.Method == http.MethodGet {
			// We only check for query parameters if the inputStruct is not nil
			if inputStruct != nil {
				if err := c.ShouldBindQuery(inputStruct); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					c.Abort()
					return
				}
			}
		}

		if inputStruct != nil {
			cleanUpInputValues(inputStruct)
		}

		if err := validate.Struct(inputStruct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Validation failed. Please check the provided data.",
				"error":   formatValidationError(err.(validator.ValidationErrors)),
			})
			c.Abort()
			return
		}
		// i have issue when not set query, it always take from prev query
		c.Set("validatedData", inputStruct)
		c.Header("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		c.Next()
	}
}

func formatValidationError(errs validator.ValidationErrors) []map[string]interface{} {
	var errorList []map[string]interface{}

	for _, e := range errs {
		errorItem := map[string]interface{}{
			"field":   e.Field(),
			"tag":     e.Tag(),
			"message": msgForTag(e),
		}
		errorList = append(errorList, errorItem)
	}

	return errorList
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "min":
		return fmt.Sprintf("This field must be at least %s characters long", fe.Param())
	case "max":
		return fmt.Sprintf("This field must be at most %s characters long", fe.Param())
	}
	return fe.Translate(nil)
}

func cleanUpInputValues(inputStruct interface{}) {
	val := reflect.ValueOf(inputStruct).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := field.Type()

		// Check if the field is a pointer and if it's nil, set it to nil
		if field.Kind() == reflect.Ptr && field.IsNil() {
			continue
		}

		// Check if the field is a pointer to int and if it's zero, set it to nil
		if fieldType.Kind() == reflect.Ptr && fieldType.Elem().Kind() == reflect.Int {
			if field.IsNil() || field.Elem().Interface() == 0 {
				field.Set(reflect.Zero(fieldType))
			}
		}
	}
}
