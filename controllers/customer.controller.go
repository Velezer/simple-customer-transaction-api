package controllers

import (
	"net/http"
	"sctrans/httperror"
	"sctrans/models"
	"sctrans/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController struct{}

type CreateCustomerInput struct {
	Name string `json:"name" binding:"required"`
}

// Register godoc
// @Summary      CreateCustomer.
// @Description  CreateCustomer.
// @Tags         Customer
// @Param        Body  body  CreateCustomerInput  true  "the body to CreateCustomer"
// @Produce      json
// @Success      200  {object}  models._Res{data=map[string]string}
// @Success      400  {object}  models._Err
// @Router       /customers [post]
func (a CustomerController) CreateCustomer(c *gin.Context) {
	var input CreateCustomerInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	if _, err := services.CustomerService.FindByName(input.Name); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusConflict))
		return
	}

	u := models.Customer{}
	u.CustomerName = input.Name

	savedCustomer, err := services.CustomerService.Save(&u)
	if err != nil {
		c.Error(err)
		return
	}

	Customer := map[string]string{
		"name": savedCustomer.CustomerName,
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": Customer})

}

type AddAddressInput struct {
	Address string `json:"address" binding:"required"`
}

// Register godoc
// @Summary      AddAddressCustomer.
// @Description  AddAddressCustomer.
// @Tags         Customer
// @Param        Body  body  AddAddressInput  true  "the body to AddAddressCustomer"
// @Produce      json
// @Success      200  {object}  models._Res{data=map[string]string}
// @Success      400  {object}  models._Err
// @Router       /customers/{id} [post]
func (a CustomerController) AddAddressCustomer(c *gin.Context) {
	var input AddAddressInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	cid, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	err = services.CustomerService.AddAddress(uint(cid), input.Address)
	if err != nil {
		c.Error(err)
		return
	}

	Customer := map[string]interface{}{
		"customer_id": cid,
		"address":     input.Address,
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": Customer})

}

// GetCustomers godoc
// @Summary      get customers
// @Description  get customers
// @Tags         Customer
// @Produce      json
// @Success      200  {object}  models._Res{data=[]models.Customer}
// @Success      500  {object}  models._Err
// @Router       /customers [get]
func (h CustomerController) GetCustomers(c *gin.Context) {
	customers, err := services.CustomerService.FindAll()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": customers})
}
