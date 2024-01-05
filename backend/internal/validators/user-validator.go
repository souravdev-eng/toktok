// validateAndCreateUser validates the user and creates a new user.
package validators

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// handleValidationError handles the validation error response.
func HandleValidationError(ctx *gin.Context, err error) {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var errors []string
	for _, ve := range validationErrors {
		field := ve.Field()
		tag := ve.Tag()

		switch field {
		case "Email":
			switch tag {
			case "required":
				errors = append(errors, "Email is required")
			case "email":
				errors = append(errors, "Invalid email format")
			}
		case "Password":
			switch tag {
			case "required":
				errors = append(errors, "Password is required")
			case "min":
				errors = append(errors, "Password must be at least 6 characters long")
			}
		default:
			errors = append(errors, fmt.Sprintf("%s %s", field, tag))
		}

	}

	ctx.JSON(http.StatusBadRequest, gin.H{"error": errors})
}
