package auth

import (
	authRepo "bookstore/repositories/auth"

	"github.com/gin-gonic/gin"
)

type authController struct {
	authRepo authRepo.AuthRepo
}

type AuthController interface {
	AuthRegister(c *gin.Context)
	AuthLogin(c *gin.Context)
}

func NewAuthController(authRepo authRepo.AuthRepo) AuthController {
	return &authController{
		authRepo: authRepo,
	}
}
