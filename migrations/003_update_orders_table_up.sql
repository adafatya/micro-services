ALTER TABLE orders
    ADD COLUMN paypal_order_id varchar(36) UNIQUE NOT NULL;