package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taco-tortilla/jwt-go/request"
	"github.com/taco-tortilla/jwt-go/service"
)

func SignUp(c *gin.Context) {
	var body request.SingUpBody

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "fialed to read body",
		})
		return
	}

	if err := service.SignUp(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	var body request.SingUpBody

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fialed to read body",
		})
		return
	}

	tokenString, err := service.Login(body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fialed to read body",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})

}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
