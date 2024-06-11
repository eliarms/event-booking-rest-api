package routes

import (
	"eliarms.events.com/models"
	"eliarms.events.com/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signup(c *gin.Context) {

	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Could not passed request Data": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save User."})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created!", "user": user})

}

func login(c *gin.Context) {

	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Could not passed request Data": err.Error()})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Logged in Successfully", "token": token})

}
