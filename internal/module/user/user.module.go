package user

import (
	"learn-go/internal/database/repository"
)

func UserServiceInit(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func UserControllerInit(userService UserService) *UserControllerImpl {
	return &UserControllerImpl{
		userService: userService,
	}
}
