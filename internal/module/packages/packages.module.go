package packages

import "learn-go/internal/database/repository"

func PackageServiceInit(repository repository.PackageRepository) *PackageServiceImpl {
	return &PackageServiceImpl{
		repository: repository,
	}
}

func PackageControllerInit(service PackageService) *PackageControllerImpl {
	return &PackageControllerImpl{
		service: service,
	}
}
