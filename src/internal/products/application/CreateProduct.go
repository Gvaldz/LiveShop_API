package application

import (
	"liveshop_api/src/internal/products/domain"
	"liveshop_api/src/internal/products/domain/entities"
)

type CreateProduct struct {
	repo domain.ProductRepository
}

func NewCreateProduct(repo domain.ProductRepository) *CreateProduct {
	return &CreateProduct{repo: repo}
}
func (cp *CreateProduct) Execute(product entities.Product) error {
	return cp.repo.Save(product)
}

// --- Delete ---
type DeleteProduct struct {
	repo domain.ProductRepository
}

func NewDeleteProduct(repo domain.ProductRepository) *DeleteProduct {
	return &DeleteProduct{repo: repo}
}
func (dp *DeleteProduct) Execute(id int32, sellerID int32) error {
	return dp.repo.Delete(id, sellerID)
}

type ListProducts struct {
	repo domain.ProductRepository
}

func NewListProducts(repo domain.ProductRepository) *ListProducts {
	return &ListProducts{repo: repo}
}
func (lp *ListProducts) Execute(sellerID int32) ([]entities.Product, error) {
	return lp.repo.GetAll(sellerID)
}

type GetProductById struct {
	repo domain.ProductRepository
}

func NewGetProductById(repo domain.ProductRepository) *GetProductById {
	return &GetProductById{repo: repo}
}
func (g *GetProductById) Execute(id int32, sellerID int32) (entities.Product, error) {
	return g.repo.GetById(id, sellerID)
}

// --- Update ---
type UpdateProduct struct {
	repo domain.ProductRepository
}

func NewUpdateProduct(repo domain.ProductRepository) *UpdateProduct {
	return &UpdateProduct{repo: repo}
}
func (up *UpdateProduct) Execute(product entities.Product) error {
	return up.repo.Update(product)
}

