package auth

import (
	"bookstore/models/db"
	"bookstore/models/request"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type ResponseLogin struct {
	Token string
}

func (ac *authController) AuthRegister(c *gin.Context) {

	var (
		body = new(request.Register)
	)

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
		})

		return
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash Password",
		})

		return
	}

	dataRegister := db.User{
		Email:    body.Email,
		Password: string(hash),
	}

	result := ac.authRepo.Register(dataRegister)

	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusOK,
		"message": "success",
	})
}

func (ac *authController) AuthLogin(c *gin.Context) {

	var (
		body     = new(request.Register)
		response = ResponseLogin{}
	)

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed",
		})
		return
	}

	user, _ := ac.authRepo.CheckEmail(body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed",
		})
		return
	}

	// Generate a JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Failed to create token",
		})
		return
	}

	ac.authRepo.SaveToken(user.Email, tokenString)
	response.Token = tokenString

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    response,
	})
}
