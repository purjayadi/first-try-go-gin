package resource

import (
	"learn-go/internal/constant"
	"learn-go/internal/database/model"
	"learn-go/internal/database/repository"
	"learn-go/pkg/dto"
	pkg "learn-go/pkg/panic"
	"net/http"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// interface
type ResourceService interface {
	AddNewResource(c *gin.Context, input *dto.ResourceDto)
	UpdateResource(c *gin.Context, input *dto.UpdateResourceDto)
	GetAllResource(c *gin.Context, userInput *dto.SearchResourceDto)
	GetResourceById(c *gin.Context, id uuid.UUID)
	DeleteResource(c *gin.Context, id uuid.UUID)
}

// resource struct
type ResourceServiceImpl struct {
	repository repository.ResourceRepository
}

// find all resources
func (u ResourceServiceImpl) GetAllResource(c *gin.Context, userInput *dto.SearchResourceDto) {
	defer pkg.PanicHandler(c)
	log.Info("start to execute get all data resource")
	log.Debug("userInput: ", userInput)
	data, err := u.repository.FindAll(userInput)
	if err != nil {
		log.Error("error while getting all data resources: ", err)
		pkg.PanicException(constant.UnknownError)
	}
	// return all data user
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "Get resources successfully", data))
}

// get resource by id
func (u ResourceServiceImpl) GetResourceById(c *gin.Context, id uuid.UUID) {
	defer pkg.PanicHandler(c)
	log.Info("start to execute get data resource by id", id)
	// get data user by id
	data, err := u.repository.FindOne(id)
	if err != nil {
		log.Error("error while getting data resource by id: ", err)
		pkg.PanicException(constant.DataNotFound)
	}
	// return data user
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "Get resource by id successfully", data))
}

func (u ResourceServiceImpl) AddNewResource(c *gin.Context, input *dto.ResourceDto) {
	defer pkg.PanicHandler(c)
	log.Info("start to execute create data resource", input)
	// create new data user
	newResource := &model.Resource{
		Name: input.Name,
	}
	err := u.repository.Create(newResource)
	if err != nil {
		log.Error("error while creating new data resource: ", err)
		pkg.PanicException(constant.UnknownError)
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "Create resource successfully", newResource))
}

// update resource
func (u ResourceServiceImpl) UpdateResource(c *gin.Context, input *dto.UpdateResourceDto) {
	defer pkg.PanicHandler(c)
	log.Info("start to execute update data resource", input)
	// check if resource exist
	resource, err := u.CheckIfResourceExist(c, input.ID)
	resource.Name = input.Name
	err = u.repository.Update(&resource)
	if err != nil {
		log.Error("error while updating data resource: ", err)
		pkg.PanicException(constant.UnknownError)
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "Update resource successfully", resource))
}

// delete resource
func (u ResourceServiceImpl) DeleteResource(c *gin.Context, id uuid.UUID) {
	defer pkg.PanicHandler(c)
	log.Info("start to execute delete data resource", id)
	// check if resource exist
	resource, err := u.CheckIfResourceExist(c, id)
	err = u.repository.Delete(&resource)
	if err != nil {
		log.Error("error while deleting data resource: ", err)
		pkg.PanicException(constant.UnknownError)
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "Delete resource successfully", resource))
}

// check if resource exist
func (u ResourceServiceImpl) CheckIfResourceExist(c *gin.Context, id uuid.UUID) (model.Resource, error) {
	defer pkg.PanicHandler(c)
	log.Info("start to execute check if resource exist")
	// check if resource exist
	resourceData, err := u.repository.FindOne(id)
	if err != nil {
		log.Warn("error while checking if resource exist: ", err)
		pkg.PanicException(constant.DataNotFound)
	}
	// return resource
	return resourceData, nil
}
