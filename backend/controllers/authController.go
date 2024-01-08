package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/souravdev-eng/toktok-api/db"
	helper "github.com/souravdev-eng/toktok-api/helpers"
	"github.com/souravdev-eng/toktok-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = db.OpenCollection(db.Client, "user")

var validate = validator.New()

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var error []string
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User

		//1. Convert the JSON data coming from clients
		if err := c.BindJSON(&user); err != nil {
			error = append(error, err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": error})
			return
		}

		//2. validate the data based on user struct
		if validateErr := validate.Struct(user); validateErr != nil {

			for _, err := range validateErr.(validator.ValidationErrors) {
				error = append(error, getErrorMessage(err))
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": error})
			return
		}

		//3. you'll check if the email has already been used by another user
		emailExistCount, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			error = append(error, "Oops! Error occurred while checking for the email")
			c.JSON(http.StatusBadRequest, gin.H{"error": error})
			return
		}

		if emailExistCount > 0 {
			error = append(error, "Oops! Email already exist. Please try a different one.")
			c.JSON(http.StatusBadRequest, gin.H{"error": error})
			return
		}
		// 4. Hash the password
		password := helper.HashPassword(*user.Password)

		// 5. Update the user password to the encrypt
		user.Password = &password

		// 6. Create extra details for user object
		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()

		// 7. Generate token and refresh token
		token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.Name, user.User_id)
		user.Token = &token
		user.Refresh_Token = &refreshToken

		// 8. If all ok, then insert into db
		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			error = append(error, fmt.Sprintf("User item was not created"))
			c.JSON(http.StatusInternalServerError, gin.H{"error": error})
			return
		}

		//9. Return status OK and send the result back
		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return err.Field() + " is required"
	case "email":
		return err.Field() + " must be a valid email address"
	case "min":
		return err.Field() + " must be at least " + err.Param() + " characters long"
	default:
		return err.Error()
	}
}
