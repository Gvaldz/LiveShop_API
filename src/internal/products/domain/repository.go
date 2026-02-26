package domain

import "liveshop_api/src/internal/products/domain/entities"

type ProductRepository interface {
    Save(product entities.Product) error
    Update(product entities.Product) error
    GetAll(sellerID int32) ([]entities.Product, error)
    GetById(id int32, sellerID int32) (entities.Product, error)
    Delete(id int32, sellerID int32) error
}