package repository

import (
	"database/sql"
	retail "testTask_retail"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type RetailPostgres struct {
	db *sqlx.DB
}

func NewRetailPostgres(db *sqlx.DB) *RetailPostgres {
	return &RetailPostgres{db: db}
}

func (r *RetailPostgres) GetOrderDetails() ([]retail.OrderDetails, error) {
	query := `
		SELECT oi.order_id, oi.product_id, p.name AS product_name, ps.shelf_id, oi.count AS order_count, 
			s.location AS shelf_location, array_agg(s2.location) AS additional_shelves
		FROM OrderItems oi
		INNER JOIN Products p ON oi.product_id = p.id
		INNER JOIN ProductShelves ps ON oi.product_id = ps.product_id
		INNER JOIN Shelves s ON ps.shelf_id = s.id
		LEFT JOIN ProductShelves ps2 ON oi.product_id = ps2.product_id AND ps2.is_main = FALSE
		LEFT JOIN Shelves s2 ON ps2.shelf_id = s2.id
		WHERE ps.is_main = TRUE
		GROUP BY oi.order_id, oi.product_id, p.name, ps.shelf_id, oi.count, s.location
		ORDER BY 
			CASE 
				WHEN s.location = 'Стеллаж А' THEN oi.order_id 
				ELSE oi.product_id 
			END
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderDetailsList []retail.OrderDetails

	for rows.Next() {
		var orderDetail retail.OrderDetails
		var additionalShelves []sql.NullString
		err := rows.Scan(&orderDetail.OrderId, &orderDetail.ProductId, &orderDetail.ProductName,
			&orderDetail.ShelfId, &orderDetail.OrderCount, &orderDetail.ShelfLocation, pq.Array(&additionalShelves))
		if err != nil {
			return nil, err
		}

		for _, shelf := range additionalShelves {
			if shelf.Valid {
				orderDetail.AdditionalShelves = append(orderDetail.AdditionalShelves, shelf.String)
			}
		}

		orderDetailsList = append(orderDetailsList, orderDetail)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orderDetailsList, nil
}
