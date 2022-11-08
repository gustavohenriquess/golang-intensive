package database

import (
	"database/sql"

	"github.com/gustavohenriquess/golang-intensive/internal/order/entity"
)

type orderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *orderRepository {
	return &orderRepository{Db: db}
}

func (r *orderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}
