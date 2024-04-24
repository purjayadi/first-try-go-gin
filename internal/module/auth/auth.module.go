package auth

import "learn-go/internal/database/repository"

// initialize auth service
func AuthServiceInit(userRepository repository.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepository: userRepository,
	}
}

// initialize auth controller
func AuthControllerInit(authService AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		authService: authService,
	}
}
