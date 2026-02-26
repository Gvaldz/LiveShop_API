package controllers

import (
	"liveshop_api/src/internal/products/application"
	"liveshop_api/src/internal/products/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductControllers struct {
	create *application.CreateProduct
	list   *application.ListProducts
	get    *application.GetProductById
	update *application.UpdateProduct
	delete *application.DeleteProduct
}

func NewProductControllers(c *application.CreateProduct, l *application.ListProducts, g *application.GetProductById, u *application.UpdateProduct, d *application.DeleteProduct) *ProductControllers {
	return &ProductControllers{create: c, list: l, get: g, update: u, delete: d}
}

type ProductRequest struct {
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
	Stock int32   `json:"stock"` 
}

func (h *ProductControllers) Create(c *gin.Context) {
	userIDInterface, _ := c.Get("userID")
	sellerID := userIDInterface.(int32)

	var req ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := entities.Product{
		Name:     req.Name,
		Price:    req.Price,
		Stock:    req.Stock,
		SellerID: sellerID,
	}

	if err := h.create.Execute(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Producto creado", "product": product})
}

func (h *ProductControllers) GetAll(c *gin.Context) {
	userIDInterface, _ := c.Get("userID")
	sellerID := userIDInterface.(int32)

	products, err := h.list.Execute(sellerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductControllers) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	userIDInterface, _ := c.Get("userID")
	sellerID := userIDInterface.(int32)

	product, err := h.get.Execute(int32(id), sellerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductControllers) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	userIDInterface, _ := c.Get("userID")
	sellerID := userIDInterface.(int32)

	var req ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := entities.Product{
		IdProduct: int32(id),
		SellerID:  sellerID,
		Name:      req.Name,
		Price:     req.Price,
		Stock:     req.Stock,
	}

	if err = h.update.Execute(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Producto actualizado", "product": product})
}

func (h *ProductControllers) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	userIDInterface, _ := c.Get("userID")
	sellerID := userIDInterface.(int32)

	if err = h.delete.Execute(int32(id), sellerID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado correctamente"})
}
