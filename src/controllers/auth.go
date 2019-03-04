package controllers

import (
	"net/http"
	"time"

	models "../models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

// Login generates a JWT token
// @Summary Login generates a JWT token
// @Description Generates a JWT Token for use with Authorized endpoints
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Success 200 {object} models.AuthToken
// @Failure 401
// @Router /auth/login [post]
func (h AuthController) Login(c *gin.Context) {

	expiresAt := time.Now().Add(time.Minute * 1).Unix()

	token := jwt.New(jwt.SigningMethodHS256)

	tokenString, err := token.SignedString([]byte("A14E45A7-D02B-4ADA-94BC-66DCBFD3181E"))
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, models.AuthToken{
		Token:     tokenString,
		TokenType: "Bearer",
		ExpiresIn: expiresAt,
	})
}
