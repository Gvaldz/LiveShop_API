package entities

type Order struct {
	IdOrder     int32
	BuyerID     int32
	ProductID   int32
	Quantity    int32
	IsDelivered bool
	CreatedAt   string
}

func NewOrder(idOrder int32, buyerID int32, productID int32, quantity int32, isDelivered bool) *Order {
	return &Order{
		IdOrder:     idOrder,
		BuyerID:     buyerID,
		ProductID:   productID,
		Quantity:    quantity,
		IsDelivered: isDelivered,
	}
}