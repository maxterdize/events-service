package routes

import (
	"events-service/models"
	"events-service/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SignUp	godoc
// @Summary		Create a User
// @Description	Create a new user and save them to the DB
// @Tags         Users
// @Param		event body models.User true "Create User"
// @Produce		application/json
// @User		user
// @Success		201
// @Router		/signup	[post]
func signup(context *gin.Context) {
	var user models.User

	err := context.BindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// Login	godoc
// @Summary		Login a User
// @Description	Login user by validating user credentials
// @Tags         Users
// @Param		event body models.User true "Login User"
// @Produce		application/json
// @User		user
// @Success		200
// @Router		/login	[post]
func login(context *gin.Context) {
	var user models.User

	err := context.BindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}

	err = user.Authenticate()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user", "error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login Successful!", "token": token})
}
