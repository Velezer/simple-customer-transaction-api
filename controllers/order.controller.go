package controllers

import (
	"net/http"
	"sctrans/httperror"
	"sctrans/models"
	"sctrans/services"
	"sync"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
}

type OrderInput struct {
	Name              string `json:"name" binding:"required"`
	CustomerAddressID uint   `json:"customer_address_id" binding:"required"`
	ProductIDs        []uint `json:"product_id" binding:"required"`
	PaymentMethodIDs  []uint `json:"payment_method_id" binding:"required"`
}

// CreateOrder godoc
// @Summary      Create Order, user role must be MERCHANT
// @Description  create Order
// @Tags         Order
// @Param        Body  body  OrderInput  true  "the body to create a Order"
// @Produce      json
// @Success      200  {object}  models._Res{data=models.Order}
// @Success      400  {object}  models._Err
// @Success      500  {object}  models._Err
// @Router       /orders [post]
func (h OrderController) CreateOrder(c *gin.Context) {
	var input OrderInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	m := models.Order{}
	m.CustomerAddressID = input.CustomerAddressID
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for _, pid := range input.ProductIDs {
			p, _ := services.ProductService.FindById(pid)
			m.Products = append(m.Products, *p)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for _, pmid := range input.PaymentMethodIDs {
			p, _ := services.PaymentMethodService.FindById(pmid)
			m.PaymentMethods = append(m.PaymentMethods, *p)
		}
		wg.Done()
	}()
	wg.Wait()

	savedOrder, err := services.OrderService.Save(&m)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedOrder})
}

// GetOrders godoc
// @Summary      get Orders, anyone can access
// @Description  get Orders
// @Tags         Order
// @Produce      json
// @Success      200  {object}  models._Res{data=[]models.Order}
// @Success      500  {object}  models._Err
// @Router       /orders [get]
func (h OrderController) GetOrders(c *gin.Context) {
	Orders, err := services.OrderService.FindAll()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": Orders})
}
