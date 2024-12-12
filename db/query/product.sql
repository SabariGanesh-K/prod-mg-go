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

-- name: GetProductsByUserID :many
SELECT * FROM products
WHERE user_id = $1 ;

-- name: GetProductByProductID :one
SELECT * FROM products
WHERE id = $1 ;

-- name: AddCompressedProductImageUrlsByID :one
UPDATE products
SET
compressed_product_images_urls = $1 
WHERE id = $2 
RETURNING *;