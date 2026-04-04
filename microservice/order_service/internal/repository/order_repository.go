package repository

import "database/sql"

type Order struct {
	ID      int32
	UserID  int32
	Product string
}


type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) CreateOrder(userID int32, product string) (*Order, error) {

	var id int32

	err := r.DB.QueryRow(
		"INSERT INTO orders(user_id,product) VALUES($1,$2) RETURNING id",
		userID, product,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &Order{ID: id, UserID: userID, Product: product}, nil

}

func (r *OrderRepository) GetOrder(id int32) (*Order, error) {

	row := r.DB.QueryRow("SELECT id,user_id,product from orders where id=$1", id)

	order := &Order{}

	err := row.Scan(&order.ID, &order.UserID, &order.Product)
	if err != nil {
		return nil, err
	}

	return order, nil

}
