-- name: CreateProduct :one
INSERT INTO products (
    id,
    user_id,
    product_name,
    product_description,
    product_price,
    product_urls,
    compressed_product_images_urls
) VALUES (
    $1,$2,$3,$4,$5,$6,'{}'
) RETURNING *;


-- name: GetProductByProductID :one
SELECT * FROM products
WHERE id = $1 ;

-- name: AddCompressedProductImageUrlsByID :one
UPDATE products
SET
compressed_product_images_urls = $1 
WHERE id = $2 
RETURNING *;

-- name: GetProductsByUserID :many
SELECT * FROM products
WHERE user_id = sqlc.arg(user_id)
  AND product_price >= COALESCE(sqlc.narg(min_price), product_price)
  AND product_price <= COALESCE(sqlc.narg(max_price), product_price)
  AND product_name ILIKE '%' || COALESCE(sqlc.narg(product_name), product_name) || '%';