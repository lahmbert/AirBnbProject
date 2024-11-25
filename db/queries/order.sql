-- name: FindOrderById :one
SELECT order_id, customer_id, employee_id, order_date, 
required_date, shipped_date, ship_via, freight, ship_name,
ship_address, ship_city, ship_region, ship_postal_code, ship_country
FROM orders where order_id=$1;

-- name: FindAllOrder :many
SELECT order_id, customer_id, employee_id, order_date, 
required_date, shipped_date, ship_via, freight, ship_name,
ship_address, ship_city, ship_region, ship_postal_code, ship_country
FROM orders;


-- name: CreateOrder :one
INSERT INTO orders(
	 customer_id, employee_id, order_date, required_date, 
	shipped_date, ship_via, freight, ship_name,
	ship_address, ship_city, ship_region, ship_postal_code, ship_country)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,$10,$11,$12,$13)
	RETURNING *;

-- name: UpdateOrderShip :one
UPDATE orders
	SET ship_name=$1
	WHERE order_id=$2
	RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders
	WHERE order_id=$1
    RETURNING *;