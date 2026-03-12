package entities

type Product struct {
    IdProduct int32
    Name      string
    Price     float64
    Stock     int32
    SellerID  int32
    ImgURL    string
}

func NewProduct(idProduct int32, name string, price float64, stock int32, sellerID int32, imgurl string) *Product {
    return &Product{
        IdProduct: idProduct,
        Name:      name,
        Price:     price,
        Stock:     stock,
        SellerID:  sellerID,
        ImgURL:    imgurl,
    }
}