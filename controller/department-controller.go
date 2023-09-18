package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/haroldpacha/customer-management/cache"
	"github.com/haroldpacha/customer-management/dto"
	"github.com/haroldpacha/customer-management/entity"
	"github.com/haroldpacha/customer-management/helper"
	"github.com/haroldpacha/customer-management/service"
)

type DepartmentController interface {
	All(context *gin.Context)
	Show(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	RefreshCache(keys ...string)
}

type departmentController struct {
	departmentService service.DepartmentService
	departmentCache   cache.DepartmentCache
}

func NewDepartmentController(departmentService service.DepartmentService, departmentCache cache.DepartmentCache) DepartmentController {
	return &departmentController{
		departmentService: departmentService,
		departmentCache:   departmentCache,
	}
}

func (c *departmentController) All(context *gin.Context) {
	var departments []entity.Department = c.departmentCache.Get("all")
	if departments == nil {
		departments = c.departmentService.All()
		c.departmentCache.Set("all", departments)
	}

	res := helper.BuildValidResponse("OK", departments)
	context.JSON(http.StatusOK, res)
}

func (c *departmentController) Show(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("parameter ID must not be empty", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var arr []entity.Department = c.departmentCache.Get(strconv.FormatUint(id, 10))
	if arr == nil {
		var department entity.Department = c.departmentService.Show(id)
		if (department == entity.Department{}) {
			res := helper.BuildErrorResponse("failed to retrieve Department", "no data with given departmentID", helper.EmptyObj{})
			context.AbortWithStatusJSON(http.StatusNotFound, res)
			return
		}
		arr = append(arr, department)
		c.departmentCache.Set(strconv.FormatUint(id, 10), arr)
	}

	res := helper.BuildValidResponse("OK", arr[0])
	context.JSON(http.StatusOK, res)
}

func (c *departmentController) Insert(context *gin.Context) {
	var departmentCreateDTO dto.DepartmentCreateDTO
	errDTO := context.ShouldBind(&departmentCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result := c.departmentService.Insert(departmentCreateDTO)
	response := helper.BuildValidResponse("OK", result)
	context.JSON(http.StatusCreated, response)

	c.RefreshCache("all")
}

func (c *departmentController) Update(context *gin.Context) {
	var departmentUpdateDTO dto.DepartmentUpdateDTO
	errDTO := context.ShouldBind(&departmentUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result := c.departmentService.Update(departmentUpdateDTO)
	response := helper.BuildValidResponse("OK", result)
	context.JSON(http.StatusOK, response)

	c.RefreshCache("all", strconv.FormatUint(departmentUpdateDTO.ID, 10))
}

func (c *departmentController) Delete(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("parameter ID must not be empty", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var department entity.Department = c.departmentService.Show(id)
	if (department == entity.Department{}) {
		res := helper.BuildErrorResponse("failed to retrieve Department", "no data with given departmentID", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	c.departmentService.Delete(department)
	message := fmt.Sprintf("Department with ID %v successfuly deleted", department.ID)
	res := helper.BuildValidResponse(message, helper.EmptyObj{})
	context.JSON(http.StatusOK, res)

	c.RefreshCache("all", strconv.FormatUint(id, 10))
}

func (c *departmentController) RefreshCache(keys ...string) {
	for _, key := range keys {
		c.departmentCache.Del(key)
	}
}
