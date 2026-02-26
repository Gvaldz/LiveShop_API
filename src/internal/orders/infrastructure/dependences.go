package infrastructure

import (
	"database/sql"
	"liveshop_api/src/internal/orders/application"
	"liveshop_api/src/internal/orders/domain"
	"liveshop_api/src/internal/orders/infrastructure/controllers"
	products_infra "liveshop_api/src/internal/products/infrastructure"

	"github.com/gin-gonic/gin"
)

type OrderDependencies struct {
	DB             *sql.DB
	AuthMiddleware gin.HandlerFunc
	Notifier       domain.INotifier
}

func NewOrderDependencies(db *sql.DB, authMiddleware gin.HandlerFunc, notifier domain.INotifier) *OrderDependencies {
	return &OrderDependencies{
		DB:             db,
		AuthMiddleware: authMiddleware,
		Notifier:       notifier, 
	}
}

func (d *OrderDependencies) GetRoutes() *OrderRoutes {
	repo := NewOrderRepo(d.DB)
	productRepo := products_infra.NewProductRepo(d.DB)

	createUC := application.NewCreateOrder(repo, d.Notifier, productRepo)
	listUC := application.NewListOrders(repo)
	getByIdUC := application.NewGetOrderById(repo)
	updateUC := application.NewUpdateOrder(repo)
	deleteUC := application.NewDeleteOrder(repo)

	orderControllers := controllers.NewOrderControllers(createUC, listUC, getByIdUC, updateUC, deleteUC)

	return NewOrderRoutes(orderControllers, d.AuthMiddleware)
}
