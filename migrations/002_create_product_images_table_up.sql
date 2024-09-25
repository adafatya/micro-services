CREATE TABLE product_images (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL REFERENCES products(id),
    image_path VARCHAR(200) NOT NULL,
    is_thumbnail BOOLEAN NOT NULL
);