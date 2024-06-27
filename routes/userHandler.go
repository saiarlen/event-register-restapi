package routes

import (
	"eventapi/models"
	"eventapi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Signup saves user to database if user is signed up.
//
// @param context - to pass to api
func signup(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)

	// error is not nil error is returned
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"messsage": "User created successfully"})

}

// login validates user credentials and generates a token for the user. It is used by clients to log in to Gin
//
// @param context - Context to pass to
func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not auth user."})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login Successful!", "token": token})
}
