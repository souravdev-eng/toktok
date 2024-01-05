package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kamva/mgm/v3"
	"github.com/souravdev-eng/toktok-api/internal/validators"
	"github.com/souravdev-eng/toktok-api/models"
)

// Signup handles user registration.
func Signup(ctx *gin.Context) {
	user := &models.User{}

	if err := ctx.ShouldBindJSON(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	err := validateAndCreateUser(user)
	if err != nil {
		validators.HandleValidationError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Signup successful!",
		"user":    user,
	})
}

// Login handles user login.
func Login(ctx *gin.Context) {
	user := &models.User{}

	if err := ctx.ShouldBindJSON(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	err := validateAndLoginUser(user)
	if err != nil {
		validators.HandleValidationError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful!",
		"user":    user,
	})
}

func validateAndCreateUser(user *models.User) error {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return err
	}

	// Additional logic for creating a new user in the database using mgm
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return mgm.Coll(user).Create(user)
}

// validateAndLoginUser validates the user for login.
func validateAndLoginUser(user *models.User) error {
	validate := validator.New()
	return validate.Struct(user)
}
