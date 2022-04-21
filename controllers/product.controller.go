package controllers

import (
	"net/http"
	"sctrans/httperror"
	"sctrans/models"
	"sctrans/services"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type ProductController struct {
}

type ProductInput struct {
	Name  int             `json:"name" binding:"required"`
	Price decimal.Decimal `json:"price" binding:"required"`
}

// CreateProduct godoc
// @Summary      Create Product, user role must be MERCHANT
// @Description  create product
// @Tags         Product
// @Param        Body  body  ProductInput  true  "the body to create a Product"
// @Produce      json
// @Success      200  {object}  models._Res{data=models.Product}
// @Success      400  {object}  models._Err
// @Success      500  {object}  models._Err
// @Router       /products [post]
func (h ProductController) CreateProduct(c *gin.Context) {
	var input ProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err).SetMeta(httperror.NewMeta(http.StatusBadRequest))
		return
	}

	m := models.Product{}
	m.Name = input.Name
	m.Price = input.Price

	savedProduct, err := services.ProductService.Save(&m)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": savedProduct})
}

// GetProducts godoc
// @Summary      get products, anyone can access
// @Description  get products
// @Tags         Product
// @Produce      json
// @Success      200  {object}  models._Res{data=[]models.Product}
// @Success      500  {object}  models._Err
// @Router       /products [get]
func (h ProductController) GetProducts(c *gin.Context) {
	products, err := services.ProductService.FindAll()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": products})
}
