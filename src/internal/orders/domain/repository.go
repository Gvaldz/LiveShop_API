package domain

import "liveshop_api/src/internal/orders/domain/entities"

type IOrder interface {
	Save(order entities.Order) error
	Update(order entities.Order) error
	GetAll(buyerID int32) ([]entities.Order, error)
	GetById(id int32, buyerID int32) (entities.Order, error)
	Delete(id int32, buyerID int32) error

}