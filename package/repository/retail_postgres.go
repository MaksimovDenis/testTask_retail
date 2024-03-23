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
	SELECT 
    oi.order_id, 
    oi.product_id, 
    p.name AS product_name, 
    (
        SELECT ps.shelf_id 
        FROM ProductShelves ps 
        WHERE ps.product_id = oi.product_id
        LIMIT 1
    ) AS shelf_id, 
    oi.count AS order_count, 
    (
        SELECT s.location 
        FROM Shelves s
        WHERE s.id = (
            SELECT ps.shelf_id 
            FROM ProductShelves ps 
            WHERE ps.product_id = oi.product_id
            LIMIT 1
        )
    ) AS shelf_location, 
    (
        SELECT array_agg(s2.location) 
        FROM Shelves s2
        WHERE EXISTS (
            SELECT 1 
            FROM ProductShelves ps2 
            WHERE s2.id = ps2.shelf_id AND ps2.product_id = oi.product_id AND ps2.is_main = FALSE
        )
    ) AS additional_shelves
	FROM 
		OrderItems oi, 
		Products p
	WHERE 
		oi.product_id = p.id
	GROUP BY 
		oi.order_id, oi.product_id, p.name, oi.count
	ORDER BY 
		CASE 
			WHEN (
				SELECT s.location 
				FROM Shelves s
				WHERE s.id = (
					SELECT ps.shelf_id 
					FROM ProductShelves ps 
					WHERE ps.product_id = oi.product_id
					LIMIT 1
				)
			) = 'Стеллаж А' THEN oi.order_id 
			ELSE oi.product_id 
    END;
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
