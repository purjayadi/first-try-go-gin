package user

import (
	"learn-go/pkg/dto"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAllUser(c *gin.Context)
	GetUserByEmail(c *gin.Context)
	AddNewUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserControllerImpl struct {
	userService UserService
}

// GetUsers             godoc
// @Security BearerAuth
// @Summary      Get users array
// @Description  Responds with the list of all users as JSON.
// @Tags         users
// @Produce      json
// @Success      200  {array}  model.User
// @Param        page  query      string  false  "page"
// @Param        pageSize  query      string  false  "pageSize"
// @Router       /user [get]
func (u UserControllerImpl) GetAllUser(c *gin.Context) {
	u.userService.GetAllUser(c, c.MustGet("validatedData").(*dto.GetUserDto))
}

// find user by email
// @Security BearerAuth
// @Summary      Get user by email
// @Description  Find user by email
// @Tags         users
// @Produce      json
// @Param        email  path      string  true  "email"
// @Success      200    {object}  model.User
// @Router       /user/{email} [get]
func (u UserControllerImpl) GetUserByEmail(c *gin.Context) {
	u.userService.GetUserByEmail(c)
}

// create new user
// @Security BearerAuth
// @Summary      Create new user
// @Description  Create new user
// @Tags         users
// @Produce      json
// @Param        userInput  body      dto.CreateUserDto  true  "userInput"
// @Success      200        {object}  model.User
// @Router       /user [post]
func (u UserControllerImpl) AddNewUser(c *gin.Context) {
	u.userService.AddNewUser(c, c.MustGet("validatedData").(*dto.CreateUserDto))
}

// update user by id
// @Security BearerAuth
// @Summary      Update user by id
// @Description  Update user by id
// @Tags         users
// @Produce      json
// @Param        id  path      string  true  "user id"
// @Param        userInput  body      dto.UpdateUserDto  true  "userInput"
// @Success      200        {object}  model.User
// @Router       /user/{id} [patch]
func (u UserControllerImpl) UpdateUser(c *gin.Context) {
	u.userService.UpdateUser(c, c.MustGet("validatedData").(*dto.UpdateUserDto))
}

// delete user by id
// @Security BearerAuth
// @Summary      Delete user by id
// @Description  Delete user by id
// @Tags         users
// @Produce      json
// @Param        id  path      string  true  "user id"
// @Success      200  {object}  model.User
// @Router       /user/{id} [delete]
func (u UserControllerImpl) DeleteUser(c *gin.Context) {
	u.userService.DeleteUser(c)
}
