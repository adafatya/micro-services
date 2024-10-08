CREATE TABLE order_products(
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL REFERENCES orders(id),
    product_id INTEGER NOT NULL REFERENCES products(id),
    quantity SMALLINT NOT NULL CHECK (quantity > 0)
);