CREATE TABLE products(
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    price INTEGER NOT NULL CHECK (price >= 0),
    quantity SMALLINT NOT NULL CHECK (quantity >= 0)
);