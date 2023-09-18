package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/haroldpacha/customer-management/dto"
	"github.com/haroldpacha/customer-management/entity"
	"github.com/haroldpacha/customer-management/helper"
	"github.com/haroldpacha/customer-management/service"
)

type CustomerController interface {
	All(context *gin.Context)
	Get(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type customerController struct {
	customerService service.CustomerService
	jwtService      service.JWTService
}

func NewCustomerController(customerService service.CustomerService, jwtService service.JWTService) CustomerController {
	return &customerController{
		customerService: customerService,
		jwtService:      jwtService,
	}
}

func (c *customerController) All(context *gin.Context) {
	var customers []entity.Customer = c.customerService.All()
	res := helper.BuildValidResponse("OK", customers)
	context.JSON(http.StatusOK, res)
}

func (c *customerController) Get(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("parameter ID must not be empty", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var customer entity.Customer = c.customerService.Get(id)
	if (customer == entity.Customer{}) {
		res := helper.BuildErrorResponse("failed to retrieve Customer", "no data with given customerID", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	res := helper.BuildValidResponse("OK", customer)
	context.JSON(http.StatusOK, res)
}

func (c *customerController) Insert(context *gin.Context) {
	var customerCreateDTO dto.CustomerCreateDTO
	errDTO := context.ShouldBind(&customerCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result := c.customerService.Insert(customerCreateDTO)
	response := helper.BuildValidResponse("OK", result)
	context.JSON(http.StatusCreated, response)
}

func (c *customerController) Update(context *gin.Context) {
	var customerUpdateDTO dto.CustomerUpdateDTO
	errDTO := context.ShouldBind(&customerUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result := c.customerService.Update(customerUpdateDTO)
	response := helper.BuildValidResponse("OK", result)
	context.JSON(http.StatusOK, response)
}

func (c *customerController) Delete(context *gin.Context) {
	var customer entity.Customer
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("parameter ID must not be empty", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	customer.ID = id
	c.customerService.Delete(customer)
	message := fmt.Sprintf("Customer with ID %v successfuly deleted", customer.ID)
	res := helper.BuildValidResponse(message, helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}
