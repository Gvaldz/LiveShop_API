package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"liveshop_api/src/internal/orders/domain/entities"
)

type orderRepo struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) *orderRepo {
	return &orderRepo{db: db}
}

func (r *orderRepo) Save(order entities.Order) error {

	query := "INSERT INTO orders (buyer_id, product_id, quantity, is_delivered) VALUES (?, ?, ?, ?)"
	_, err := r.db.Exec(query, order.BuyerID, order.ProductID, order.Quantity, order.IsDelivered)
	if err != nil {
		return fmt.Errorf("error al guardar pedido: %w", err)
	}
	return nil
}

func (r *orderRepo) GetAll(buyerID int32) ([]entities.Order, error) {
	query := "SELECT id, product_id, quantity, is_delivered, created_at FROM orders WHERE buyer_id = ?"

	rows, err := r.db.Query(query, buyerID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener pedidos: %w", err)
	}
	defer rows.Close()

	var orders []entities.Order
	for rows.Next() {
		var order entities.Order
		order.BuyerID = buyerID 
		if err := rows.Scan(&order.IdOrder, &order.ProductID, &order.Quantity, &order.IsDelivered, &order.CreatedAt); err != nil {
			return nil, fmt.Errorf("error al escanear pedido: %w", err)
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *orderRepo) GetById(id int32, buyerID int32) (entities.Order, error) {
	var order entities.Order
	query := "SELECT id, product_id, quantity, is_delivered, created_at FROM orders WHERE id = ? AND buyer_id = ?"

	order.BuyerID = buyerID
	err := r.db.QueryRow(query, id, buyerID).Scan(&order.IdOrder, &order.ProductID, &order.Quantity, &order.IsDelivered, &order.CreatedAt)
	if err != nil {
		return order, fmt.Errorf("pedido no encontrado o acceso denegado: %w", err)
	}
	return order, nil
}

func (r *orderRepo) Update(order entities.Order) error {
	query := "UPDATE orders SET is_delivered=?, quantity=? WHERE id = ? AND buyer_id = ?"

	result, err := r.db.Exec(query, order.IsDelivered, order.Quantity, order.IdOrder, order.BuyerID)
	if err != nil {
		return fmt.Errorf("error al actualizar pedido: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil || rows == 0 {
		return fmt.Errorf("pedido no encontrado o no tienes permiso para editarlo")
	}

	return nil
}

func (r *orderRepo) Delete(idOrder int32, buyerID int32) error {
	query := "DELETE FROM orders WHERE id = ? AND buyer_id = ?"

	result, err := r.db.Exec(query, idOrder, buyerID)
	if err != nil {
		return fmt.Errorf("error al eliminar pedido: %w", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("pedido no encontrado o no tienes permiso para eliminarlo")
	}

	log.Println("[orderRepo] - pedido eliminado correctamente")
	return nil
}