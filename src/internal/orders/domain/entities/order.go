package entities

type Order struct {
    IdOrder     int32  `json:"id_order"`
    BuyerID     int32  `json:"buyer_id"`
    BuyerName   string `json:"buyer_name"`   
    BuyerNumber string `json:"buyer_number"` 
    ProductID   int32  `json:"product_id"`
    Quantity    int32  `json:"quantity"`
    IsDelivered bool   `json:"is_delivered"`
    CreatedAt   string `json:"created_at"`
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