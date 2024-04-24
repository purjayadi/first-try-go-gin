package user

import (
	"learn-go/internal/constant"
	"learn-go/internal/database/model"
	"learn-go/internal/database/repository"
	"learn-go/internal/utils"
	"learn-go/pkg/dto"
	pkg "learn-go/pkg/panic"
	"net/http"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	GetAllUser(c *gin.Context, userInput *dto.GetUserDto)
	GetUserByEmail(c *gin.Context)
	AddNewUser(c *gin.Context, userInput *dto.CreateUserDto)
	UpdateUser(c *gin.Context, userInput *dto.UpdateUserDto)
	DeleteUser(c *gin.Context)
	// GetUserById(c *gin.Context)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func (u UserServiceImpl) GetAllUser(c *gin.Context, userInput *dto.GetUserDto) {
	defer pkg.PanicHandler(c)
	log.Info("start to execute get all data user", userInput)
	// get all data user
	var page, pageSize int
	if userInput.Page != nil {
		page = *userInput.Page
	}
	if userInput.PageSize != nil {
		pageSize = *userInput.PageSize
	}
	log.Debug("page: ", page, "pageSize: ", pageSize)
	data, err := u.userRepository.FindAll(page, pageSize)
	if err != nil {
		log.Error("error while getting all data user: ", err)
		pkg.PanicException(constant.UnknownError)
	}
	// return all data user
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "Get users successfully", data))
}

// find user by email
func (u UserServiceImpl) GetUserByEmail(c *gin.Context) {
	defer pkg.PanicHandler(c)
	log.Info("start to execute get user by email")
	// get email from request
	email := c.Param("email")
	// get user by email
	user, err := u.userRepository.FindOneByEmail(email)
	if err != nil {
		log.Error("error while getting user by email: ", err)
		pkg.PanicException(constant.UnknownError)
	}
	// return user
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "Get user by email successfully", user))
}

// create new user
func (u UserServiceImpl) AddNewUser(c *gin.Context, userInput *dto.CreateUserDto) {
	defer pkg.PanicHandler(c)
	log.Info("start to execute add new user")
	// check if email already exist
	_, err := u.userRepository.FindOneByEmail(userInput.Email)
	if err == nil {
		log.Error("error while creating new user: ", err)
		pkg.PanicException(constant.Conflict, "Email already exist")
	}
	// hash password
	hashedPassword, err := utils.HashPassword(userInput.Password)
	// create new user with hashed password
	user := model.User{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: string(hashedPassword),
	}
	err = u.userRepository.Create(&user)
	// return user
	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, "Create new user successfully", user))
}

// update user
func (u UserServiceImpl) UpdateUser(c *gin.Context, userInput *dto.UpdateUserDto) {
	defer pkg.PanicHandler(c)
	log.Info("start to execute update user")
	// get id from request
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Error("Invalid user ID: ", err)
		pkg.PanicException(constant.InvalidRequest, "Invalid user ID")
	}
	// get user by id
	user, err := u.userRepository.FindOne(id)
	if err != nil {
		log.Error("error while getting user: ", err)
		pkg.PanicException(constant.UnknownError)
	}
	// hash password
	if userInput.Password != "" {
		// hash password and put on var hashedPassword
		hashedPassword, _ := utils.HashPassword(userInput.Password)
		user.Password = string(hashedPassword)
	}
	// update user with hashed password
	user.Name = userInput.Name
	user.Email = userInput.Email
	err = u.userRepository.Update(&user)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "Update user successfully", user))
}

// delete user
func (u UserServiceImpl) DeleteUser(c *gin.Context) {
	defer pkg.PanicHandler(c)
	log.Info("start to execute delete user")
	// get id from request
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Error("Invalid user ID: ", err)
		pkg.PanicException(constant.InvalidRequest, "Invalid user ID")
	}
	// get user by id
	user, err := u.userRepository.FindOne(id)
	if err != nil {
		log.Error("error while getting user: ", err)
		pkg.PanicException(constant.DataNotFound)
	}
	// delete user
	err = u.userRepository.Delete(&user)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "Delete user successfully", user))
}
