package infrastructure

import (
	"database/sql"
	"liveshop_api/src/internal/orders/application"
	"liveshop_api/src/internal/orders/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

type OrderDependencies struct {
	DB             *sql.DB
	AuthMiddleware gin.HandlerFunc
}

func NewOrderDependencies(db *sql.DB, authMiddleware gin.HandlerFunc) *OrderDependencies {
	return &OrderDependencies{
		DB:             db,
		AuthMiddleware: authMiddleware,
	}
}

func (d *OrderDependencies) GetRoutes() *OrderRoutes {
	// 1. Repo
	repo := NewOrderRepo(d.DB)

	// 2. Casos de Uso
	createUC := application.NewCreateOrder(repo)
	listUC := application.NewListOrders(repo)
	getByIdUC := application.NewGetOrderById(repo)
	updateUC := application.NewUpdateOrder(repo)
	deleteUC := application.NewDeleteOrder(repo)

	// 3. Controladores
	orderControllers := controllers.NewOrderControllers(createUC, listUC, getByIdUC, updateUC, deleteUC)

	// 4. Rutas
	return NewOrderRoutes(orderControllers, d.AuthMiddleware)
}
