package application

import (
	"liveshop_api/src/internal/orders/domain"
	"liveshop_api/src/internal/orders/domain/entities"
)

// --- Create ---
type CreateOrder struct {
	repo domain.IOrder
}
func NewCreateOrder(repo domain.IOrder) *CreateOrder { return &CreateOrder{repo: repo} }
func (co *CreateOrder) Execute(order entities.Order) error { return co.repo.Save(order) }

// --- Delete ---
type DeleteOrder struct {
	repo domain.IOrder
}
func NewDeleteOrder(repo domain.IOrder) *DeleteOrder { return &DeleteOrder{repo: repo} }
func (do *DeleteOrder) Execute(id int32, buyerID int32) error { return do.repo.Delete(id, buyerID) }

// --- List All ---
type ListOrders struct {
	repo domain.IOrder
}
func NewListOrders(repo domain.IOrder) *ListOrders { return &ListOrders{repo: repo} }
func (lo *ListOrders) Execute(buyerID int32) ([]entities.Order, error) { return lo.repo.GetAll(buyerID) }

// --- Get By Id ---
type GetOrderById struct {
	repo domain.IOrder
}
func NewGetOrderById(repo domain.IOrder) *GetOrderById { return &GetOrderById{repo: repo} }
func (goi *GetOrderById) Execute(id int32, buyerID int32) (entities.Order, error) { return goi.repo.GetById(id, buyerID) }

// --- Update ---
type UpdateOrder struct {
	repo domain.IOrder
}
func NewUpdateOrder(repo domain.IOrder) *UpdateOrder { return &UpdateOrder{repo: repo} }
func (uo *UpdateOrder) Execute(order entities.Order) error { return uo.repo.Update(order) }