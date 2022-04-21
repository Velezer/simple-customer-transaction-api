package controllers

import (
	"net/http"
	"sctrans/httperror"
	"sctrans/models"
	"sctrans/services"

	"github.com/gin-gonic/gin"
)

type PaymentMethodController struct {
}

type PaymentMethodInput struct {
	Name string `json:"name" binding:"required"`
}

// CreatePaymentMethod godoc
// @Summary      Create PaymentMethod, user role must be MERCHANT
// @Description  create PaymentMethod
// @Tags         PaymentMethod
// @Param        Body  body  PaymentMethodInput  true  "the body to create a PaymentMethod"
// @Produce      json
// @Success      200  {object}  models._Res{data=models.PaymentMethod}
// @Success      400  {object}  models._Err
// @Success      500  {object}  models._Err
// @Router       /payment-methods [post]
func (h PaymentMethodController) CreatePaymentMethod(c *gin.Context) {
	var input PaymentMethodInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	m := models.PaymentMethod{}
	m.Name = input.Name
	m.IsActive = true

	savedPaymentMethod, err := services.PaymentMethodService.Save(&m)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedPaymentMethod})
}

// GetPaymentMethods godoc
// @Summary      get PaymentMethods, anyone can access
// @Description  get PaymentMethods
// @Tags         PaymentMethod
// @Produce      json
// @Success      200  {object}  models._Res{data=[]models.PaymentMethod}
// @Success      500  {object}  models._Err
// @Router       /payment-methods [get]
func (h PaymentMethodController) GetPaymentMethods(c *gin.Context) {
	PaymentMethods, err := services.PaymentMethodService.FindAll()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": PaymentMethods})
}
