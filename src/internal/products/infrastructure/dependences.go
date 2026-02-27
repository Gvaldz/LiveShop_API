package infrastructure

import (
	"database/sql"
	"liveshop_api/src/internal/products/application"
	"liveshop_api/src/internal/products/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

type ProductDependencies struct {
	DB             *sql.DB
	AuthMiddleware gin.HandlerFunc
}

func NewProductDependencies(db *sql.DB, authMiddleware gin.HandlerFunc) *ProductDependencies {
	return &ProductDependencies{
		DB:             db,
		AuthMiddleware: authMiddleware,
	}
}

func (d *ProductDependencies) GetRoutes() *ProductRoutes {
    repo := NewProductRepo(d.DB)

    createUC := application.NewCreateProduct(repo)
    listUC := application.NewListProducts(repo)
    getByIdUC := application.NewGetProductById(repo)
    updateUC := application.NewUpdateProduct(repo)
    deleteUC := application.NewDeleteProduct(repo)
    listPublicUC := application.NewListAllProductsPublic(repo)
	getByIdPublicUC := application.NewGetProductByIdPublic(repo)

    productControllers := controllers.NewProductControllers(createUC, listUC, getByIdUC, updateUC, deleteUC, listPublicUC, getByIdPublicUC)

    return NewProductRoutes(productControllers, d.AuthMiddleware)
}