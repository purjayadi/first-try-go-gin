package resource

import (
	"learn-go/internal/database/repository"
)

func ResourceServiceInit(repository repository.ResourceRepository) *ResourceServiceImpl {
	return &ResourceServiceImpl{
		repository: repository,
	}
}

func ResourceControllerInit(service ResourceService) *ResourceControllerImpl {
	return &ResourceControllerImpl{
		service: service,
	}
}
