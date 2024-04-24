package packages

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

type PackageService interface {
	GetAllPackage(c *gin.Context, userInput *dto.GetPackageDto)
	GetPackageById(c *gin.Context, id uuid.UUID)
	AddNewPackage(c *gin.Context, userInput *dto.PackageDto)
	UpdatePackage(c *gin.Context, userInput *dto.UpdatePackageDto)
	DeletePackage(c *gin.Context, id uuid.UUID)
}

type PackageServiceImpl struct {
	repository repository.PackageRepository
}

func (u PackageServiceImpl) GetAllPackage(c *gin.Context, userInput *dto.GetPackageDto) {
	defer pkg.PanicHandler(c)
	log.Info("start to execute get all data package", userInput)
	// get all data user
	var page, pageSize int
	if userInput.Page != nil {
		page = *userInput.Page
	}
	if userInput.PageSize != nil {
		pageSize = *userInput.PageSize
	}
	log.Debug("page: ", page, "pageSize: ", pageSize)
	data, err := u.repository.FindAll(page, pageSize)
	if err != nil {
		log.Error("error while getting all data packages: ", err)
		pkg.PanicException(constant.UnknownError)
	}
	// return all data user
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "Get packages successfully", data))
}

// find package by id
func (u PackageServiceImpl) GetPackageById(c *gin.Context, id uuid.UUID) {
	defer pkg.PanicHandler(c)
	log.Info("start to execute get package by id")
	user, err := u.repository.FindOne(id)
	if err != nil {
		log.Warn("error while getting package by id: ", err)
		pkg.PanicException(constant.DataNotFound)
	}
	// return user
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "Get package by id successfully", user))
}

// create new user
func (u PackageServiceImpl) AddNewPackage(c *gin.Context, userInput *dto.PackageDto) {
	defer pkg.PanicHandler(c)
	log.Info("start to execute add new package")

	newPackage := model.Package{
		Name:        userInput.Name,
		Description: userInput.Description,
		Price:       userInput.Price,
		Image:       userInput.Image,
	}

	err := u.repository.Create(&newPackage)
	if err != nil {
		log.Error("error while creating new package: ", err)
		pkg.PanicException(constant.UnknownError)
	}
	// return user
	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, "Create new package successfully", newPackage))
}

// update package
func (u PackageServiceImpl) UpdatePackage(c *gin.Context, userInput *dto.UpdatePackageDto) {
	defer pkg.PanicHandler(c)
	log.Info("start to execute update package", userInput)

	// check is package exist
	packages, err := u.checkPackageExist(c, userInput.ID)

	packages.Name = userInput.Name
	packages.Description = userInput.Description
	packages.Price = userInput.Price
	packages.Image = userInput.Image
	err = u.repository.Update(&packages)
	if err != nil {
		log.Error("error while updating package: ", err)
		pkg.PanicException(constant.UnknownError)
	}
	// return user
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "Update package successfully", packages))
}

// DeletePackage delete package by id
func (u PackageServiceImpl) DeletePackage(c *gin.Context, id uuid.UUID) {
	defer pkg.PanicHandler(c)
	log.Info("start to execute delete package by id")
	// get package by id
	packageData, _ := u.checkPackageExist(c, id)
	// delete package
	err := u.repository.Delete(&packageData)
	if err != nil {
		log.Error("error while deleting package: ", err)
		pkg.PanicException(constant.UnknownError)
	}
	// return package
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "Delete package successfully", packageData))
}

// check package exist and return package
func (u PackageServiceImpl) checkPackageExist(c *gin.Context, id uuid.UUID) (model.Package, error) {
	// get package by id
	packageData, err := u.repository.FindOne(id)
	if err != nil {
		log.Error("error while getting package by id: ", err)
		pkg.PanicException(constant.DataNotFound)
	}
	// return package
	return packageData, nil
}
