package resource

import (
	"learn-go/pkg/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ResourceController interface {
	GetAllResource(c *gin.Context)
	GetResourceById(c *gin.Context)
	AddNewResource(c *gin.Context)
	UpdateResource(c *gin.Context)
	DeleteResource(c *gin.Context)
}

type ResourceControllerImpl struct {
	service ResourceService
}

// get all resources	godoc
// @Security BearerAuth
// @Summary      Get resources array
// @Description  Responds with the list of all resources as JSON.
// @Tags         resource
// @Produce      json
// @Success      200  {array}  model.Resource
// @Param        page  query      string  false  "page"
// @Param        pageSize  query      string  false  "pageSize"
// @Param        search  query      string  false  "search"
// @Router       /resource [get]
func (u ResourceControllerImpl) GetAllResource(c *gin.Context) {
	// get data from request query
	filter := dto.SearchResourceDto{}
	u.service.GetAllResource(c, &filter)
}

// get resource by id		godoc
// @Security BearerAuth
// @Summary      Get resource by id
// @Description  Find resource by id
// @Tags         resource
// @Produce      json
// @Param        id  path      string  true  "id"
// @Success      200    {object}  model.Resource
// @Router       /resource/{id} [get]
func (u ResourceControllerImpl) GetResourceById(c *gin.Context) {
	// get id and convert to uuid
	id, _ := uuid.Parse(c.Param("id"))
	u.service.GetResourceById(c, id)
}

// create new resource
// @Security BearerAuth
// @Summary      Create new resource
// @Description  Create new resource
// @Tags         resource
// @Produce      json
// @Param        body  body      dto.ResourceDto  true  "body"
// @Success      200    {object}  model.Resource
// @Router       /resource [post]
func (u ResourceControllerImpl) AddNewResource(c *gin.Context) {
	u.service.AddNewResource(c, c.MustGet("validatedData").(*dto.ResourceDto))
}

// update resource godoc
// @Security BearerAuth
// @Summary      Update resource
// @Description  Update resource
// @Tags         resource
// @Produce      json
// @Param        body  body      dto.UpdateResourceDto  true  "body"
// @Success      200    {object}  model.Resource
// @Router       /resource/{id} [patch]
func (u ResourceControllerImpl) UpdateResource(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	validatedData := c.MustGet("validatedData").(*dto.UpdateResourceDto)
	validatedData.ID = id
	u.service.UpdateResource(c, validatedData)
}

// delete resource godoc
// @Security BearerAuth
// @Summary      Delete resource
// @Description  Delete resource
// @Tags         resource
// @Produce      json
// @Param        id  path      string  true  "id"
// @Success      200   {object}  model.Resource
// @Router       /resource/{id} [delete]
func (u ResourceControllerImpl) DeleteResource(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	u.service.DeleteResource(c, id)
}
