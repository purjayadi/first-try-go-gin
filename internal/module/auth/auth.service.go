package auth

import (
	"learn-go/internal/constant"
	"learn-go/internal/database/model"
	"learn-go/internal/database/repository"
	"learn-go/internal/utils"
	"learn-go/pkg/dto"
	pkg "learn-go/pkg/panic"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// interface auth
type AuthService interface {
	Login(c *gin.Context, userInput *dto.LoginDto)
	Register(c *gin.Context, userInput *dto.CreateUserDto)
}

// struct auth
type AuthServiceImpl struct {
	userRepository repository.UserRepository
}

// function login
func (a AuthServiceImpl) Login(c *gin.Context, userInput *dto.LoginDto) {
	defer pkg.PanicHandler(c)
	log.Info("start to execute login")
	// get user by email
	user, err := a.userRepository.FindOneByEmail(userInput.Email)
	if err != nil {
		pkg.PanicException(constant.Unauthorized, "Email or password is incorrect")
	}
	// compare password
	isMatch := utils.ComparePassword(user.Password, userInput.Password)
	log.Info("is match: ", isMatch)
	if isMatch != nil {
		pkg.PanicException(constant.Unauthorized, "Email or password is incorrect")
	}
	// generate token
	token, err := utils.GenerateToken(user.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	// return token
	c.JSON(http.StatusOK, gin.H{"accessToken": token, "message": "Login successfully", "status": "SUCCESS"})
}

// function register
func (a AuthServiceImpl) Register(c *gin.Context, userInput *dto.CreateUserDto) {
	defer pkg.PanicHandler(c)
	log.Info("start to execute register")
	_, err := a.userRepository.FindOneByEmail(userInput.Email)
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
	err = a.userRepository.Create(&user)
	// return user
	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, "Register successfully", nil))
}
