package infrastructure

import (
    "database/sql"
    "fmt"
    "log"
    "liveshop_api/src/internal/products/domain/entities"
)

type productRepo struct {
    db *sql.DB
}

func NewProductRepo(db *sql.DB) *productRepo {
    return &productRepo{db: db}
}

func (r *productRepo) Save(product entities.Product) error {
    query := "INSERT INTO products (seller_id, name, price, stock, img_url) VALUES (?, ?, ?, ?, ?)"
    _, err := r.db.Exec(query, product.SellerID, product.Name, product.Price, product.Stock, product.ImgURL)
    if err != nil {
        return fmt.Errorf("error al guardar producto: %w", err)
    }
    return nil
}

func (r *productRepo) GetAll(sellerID int32) ([]entities.Product, error) {
    query := "SELECT id, name, price, stock, img_url FROM products WHERE seller_id = ?"

    rows, err := r.db.Query(query, sellerID)
    if err != nil {
        return nil, fmt.Errorf("error al obtener productos: %w", err)
    }
    defer rows.Close()

    var products []entities.Product
    for rows.Next() {
        var product entities.Product
        product.SellerID = sellerID 
        if err := rows.Scan(&product.IdProduct, &product.Name, &product.Price, &product.Stock, &product.ImgURL); err != nil {
            return nil, fmt.Errorf("error al escanear producto: %w", err)
        }
        products = append(products, product)
    }
    return products, nil
}

func (r *productRepo) GetById(id int32, sellerID int32) (entities.Product, error) {
    var product entities.Product
    query := "SELECT id, name, price, stock, img_url FROM products WHERE id = ? AND seller_id = ?"

    product.SellerID = sellerID
    err := r.db.QueryRow(query, id, sellerID).Scan(&product.IdProduct, &product.Name, &product.Price, &product.Stock, &product.ImgURL)
    if err != nil {
        return product, fmt.Errorf("producto no encontrado o acceso denegado: %w", err)
    }
    return product, nil
}

func (r *productRepo) Update(product entities.Product) error {
    query := "UPDATE products SET name=?, price=?, stock=?, img_url=? WHERE id = ? AND seller_id = ?"

    result, err := r.db.Exec(query, product.Name, product.Price, product.Stock, product.ImgURL, product.IdProduct, product.SellerID,)
    if err != nil {
        return fmt.Errorf("error al actualizar producto: %w", err)
    }

    rows, err := result.RowsAffected()
    if err != nil || rows == 0 {
        return fmt.Errorf("producto no encontrado o no tienes permiso para editarlo")
    }

    return nil
} 

func (r *productRepo) Delete(idProduct int32, sellerID int32) error {
    query := "DELETE FROM products WHERE id = ? AND seller_id = ?"

    result, err := r.db.Exec(query, idProduct, sellerID)
    if err != nil {
        return fmt.Errorf("error al eliminar producto: %w", err)
    }

    rows, _ := result.RowsAffected()
    if rows == 0 {
        return fmt.Errorf("producto no encontrado o no tienes permiso para eliminarlo")
    }

    log.Println("[productRepo] - producto eliminado correctamente")
    return nil
}