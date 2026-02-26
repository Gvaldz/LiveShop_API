package controllers

import (
	"liveshop_api/src/internal/orders/application"
	"liveshop_api/src/internal/orders/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderControllers struct {
	create *application.CreateOrder
	list   *application.ListOrders
	get    *application.GetOrderById
	update *application.UpdateOrder
	delete *application.DeleteOrder
}

func NewOrderControllers(c *application.CreateOrder, l *application.ListOrders, g *application.GetOrderById, u *application.UpdateOrder, d *application.DeleteOrder) *OrderControllers {
	return &OrderControllers{create: c, list: l, get: g, update: u, delete: d}
}

type OrderRequest struct {
	ProductID   int32 `json:"product_id" binding:"required"`
	Quantity    int32 `json:"quantity" binding:"required"`
	IsDelivered bool  `json:"is_delivered"`
}

func (h *OrderControllers) Create(c *gin.Context) {
	userIDInterface, _ := c.Get("userID")
	buyerID := userIDInterface.(int32)

	var req OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := entities.Order{
		BuyerID:     buyerID,
		ProductID:   req.ProductID,
		Quantity:    req.Quantity,
		IsDelivered: false, // Por defecto al crear un pedido no está entregado
	}

	if err := h.create.Execute(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Pedido creado", "order": order})
}

func (h *OrderControllers) GetAll(c *gin.Context) {
	userIDInterface, _ := c.Get("userID")
	buyerID := userIDInterface.(int32)

	orders, err := h.list.Execute(buyerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *OrderControllers) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	userIDInterface, _ := c.Get("userID")
	buyerID := userIDInterface.(int32)

	order, err := h.get.Execute(int32(id), buyerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderControllers) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	userIDInterface, _ := c.Get("userID")
	buyerID := userIDInterface.(int32)

	var req OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := entities.Order{
		IdOrder:     int32(id),
		BuyerID:     buyerID,
		ProductID:   req.ProductID,
		Quantity:    req.Quantity,
		IsDelivered: req.IsDelivered,
	}

	if err = h.update.Execute(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pedido actualizado", "order": order})
}

func (h *OrderControllers) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	userIDInterface, _ := c.Get("userID")
	buyerID := userIDInterface.(int32)

	if err = h.delete.Execute(int32(id), buyerID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pedido eliminado correctamente"})
}
