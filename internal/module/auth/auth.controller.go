package auth

import (
	"learn-go/pkg/dto"

	"github.com/gin-gonic/gin"
)

// interface auth controller
type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

// struct auth controller
type AuthControllerImpl struct {
	authService AuthService
}

// Login             godoc
// @Summary      Login
// @Description  Login
// @Tags         auth
// @Produce      json
// @Param        userInput  body dto.LoginDto  true  "userInput"
// @Success      200    {object}  model.User
// @Router       /auth/login [post]
func (a AuthControllerImpl) Login(c *gin.Context) {
	a.authService.Login(c, c.MustGet("validatedData").(*dto.LoginDto))
}

// Register          godoc
// @Summary      Register
// @Description  Register
// @Tags         auth
// @Produce      json
// @Param        userInput  body dto.CreateUserDto  true  "userInput"
// @Success      200    {object}  model.User
// @Router       /auth/register [post]
func (a AuthControllerImpl) Register(c *gin.Context) {
	a.authService.Register(c, c.MustGet("validatedData").(*dto.CreateUserDto))
}
