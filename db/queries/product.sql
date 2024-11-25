-- name: FindProductById :one
SELECT product_id, product_name, supplier_id, 
category_id, quantity_per_unit, unit_price, 
units_in_stock, units_on_order, reorder_level, discontinued,
product_image 
FROM products WHERE product_id =$1;

-- name: FindAllProduct :many
SELECT product_id, product_name, supplier_id, 
category_id, quantity_per_unit, unit_price, 
units_in_stock, units_on_order, reorder_level, discontinued,product_image
	FROM products;

-- name: FindAllProductPaging :many
SELECT product_id, product_name, supplier_id, 
category_id, quantity_per_unit, unit_price, 
units_in_stock, units_on_order, reorder_level, discontinued,product_image
FROM products
ORDER BY product_id
LIMIT $1 OFFSET $2;	

-- name: CreateProduct :one
INSERT INTO products(
	product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued,product_image)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,$10)
	RETURNING *;

-- name: UpdateProduct :one
UPDATE products
	SET  product_name=$1, supplier_id=$2, category_id=$3, 
	quantity_per_unit=$4, unit_price=$5, units_in_stock=$6, 
	units_on_order=$7, reorder_level=$8, discontinued=$9,product_image=$10
	WHERE product_id=$11
	RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
	WHERE product_id=$1
    RETURNING *;