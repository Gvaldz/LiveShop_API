package application

import (
	"fmt"
	"liveshop_api/src/internal/orders/domain"
	"liveshop_api/src/internal/orders/domain/entities"
	productDomain "liveshop_api/src/internal/products/domain"
	userDomain "liveshop_api/src/internal/users/domain"
)

type CreateOrder struct {
    repo        domain.IOrder
    notifier    domain.INotifier
    productRepo productDomain.ProductRepository
    userRepo    userDomain.UserRepository 
}

func NewCreateOrder(r domain.IOrder, n domain.INotifier, p productDomain.ProductRepository, u userDomain.UserRepository) *CreateOrder {
    return &CreateOrder{repo: r, notifier: n, productRepo: p, userRepo: u}
}

func (co *CreateOrder) Execute(order entities.Order) error {
    if err := co.repo.Save(order); err != nil {
        return err
    }

    buyer, err := co.userRepo.GetUserByID(order.BuyerID)
    if err != nil {
        fmt.Println("Error obteniendo datos del comprador:", err)
    }

    product, err := co.productRepo.GetByIdPublic(order.ProductID)
    if err == nil {
        notificacion := map[string]interface{}{
            "type":         "NEW_ORDER",
            "product_id":   order.ProductID,
            "product_name": product.Name,
            "quantity":     order.Quantity,
            "buyer_name":   buyer.Name,   
            "buyer_number": buyer.Number, 
            "message":      "¡Tienes un nuevo pedido!",
        }
        
        go co.notifier.NotifyUser(product.SellerID, notificacion)
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