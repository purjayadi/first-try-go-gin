package packages

import (
	"learn-go/pkg/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PackageController interface {
	GetAllPackage(c *gin.Context)
	GetPackageById(c *gin.Context)
	AddNewPackage(c *gin.Context)
	UpdatePackage(c *gin.Context)
	DeletePackage(c *gin.Context)
}

type PackageControllerImpl struct {
	service PackageService
}

// getPackages             godoc
// @Security BearerAuth
// @Summary      Get packages array
// @Description  Responds with the list of all packages as JSON.
// @Tags         packages
// @Produce      json
// @Success      200  {array}  model.Package
// @Param        page  query      string  false  "page"
// @Param        pageSize  query      string  false  "pageSize"
// @Router       /package [get]
func (u PackageControllerImpl) GetAllPackage(c *gin.Context) {
	u.service.GetAllPackage(c, c.MustGet("validatedData").(*dto.GetPackageDto))
}

// find package by id		godoc
// @Security BearerAuth
// @Summary      Get package by id
// @Description  Find package by id
// @Tags         packages
// @Produce      json
// @Param        id  path      string  true  "id"
// @Success      200    {object}  model.Package
// @Router       /package/{id} [get]
func (u PackageControllerImpl) GetPackageById(c *gin.Context) {
	// get id and convert to uuid
	id, _ := uuid.Parse(c.Param("id"))
	u.service.GetPackageById(c, id)
}

// create new package		godoc
// @Security BearerAuth
// @Summary      Create new package
// @Description  Create new package
// @Tags         packages
// @Produce      json
// @Param        package  body      dto.PackageDto  true  "package"
// @Success      200    {object}  model.Package
// @Router       /package [post]
func (u PackageControllerImpl) AddNewPackage(c *gin.Context) {
	u.service.AddNewPackage(c, c.MustGet("validatedData").(*dto.PackageDto))
}

// update package		godoc
// @Security BearerAuth
// @Summary      Update package
// @Description  Update package
// @Tags         packages
// @Produce      json
// @Param        id  path      string  true  "id"
// @Param        package  body      dto.UpdatePackageDto  true  "package"
// @Success      200    {object}  model.Package
// @Router       /package/{id} [put]
func (u PackageControllerImpl) UpdatePackage(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	validatedData := c.MustGet("validatedData").(*dto.UpdatePackageDto)
	validatedData.ID = id
	u.service.UpdatePackage(c, validatedData)
}

// delete package by id		godoc
// @Security BearerAuth
// @Summary      Delete package by id
// @Description  Delete package by id
// @Tags         packages
// @Produce      json
// @Param        id  path      string  true  "id"
// @Success      200    {object}  model.Package
// @Router       /package/{id} [delete]
func (u PackageControllerImpl) DeletePackage(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	u.service.DeletePackage(c, id)
}
