package application

import (
	"fmt"
	"liveshop_api/src/internal/orders/domain"
	"liveshop_api/src/internal/orders/domain/entities"
	productDomain "liveshop_api/src/internal/products/domain"
)

type CreateOrder struct {
	repo     domain.IOrder
	notifier domain.INotifier 
    productRepo productDomain.ProductRepository 
}

func NewCreateOrder(repo domain.IOrder, notifier domain.INotifier, productRepo productDomain.ProductRepository) *CreateOrder {
	return &CreateOrder{repo: repo, notifier: notifier, productRepo: productRepo}
}

func (co *CreateOrder) Execute(order entities.Order) error {
    if err := co.repo.Save(order); err != nil {
        return err
    }

    product, err := co.productRepo.GetByIdPublic(order.ProductID)
    if err == nil {
        notificacion := map[string]interface{}{
            "type":         "NEW_ORDER",
			"buyer_name":   order.BuyerName,
			"buyer_number": order.BuyerNumber,
            "product_id":   order.ProductID,
            "product_name": product.Name, 
            "quantity":     order.Quantity,
            "message":      "¡Tienes un nuevo pedido!",
        }
        
        go co.notifier.NotifyUser(product.SellerID, notificacion)
    } else {
        fmt.Println("Error buscando producto para notificar:", err) 
    }

    return nil
}
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