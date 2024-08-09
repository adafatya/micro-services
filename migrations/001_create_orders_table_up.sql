CREATE TABLE orders(
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    user_address_id INTEGER NOT NULL REFERENCES user_addresses(id),
    total_price BIGINT NOT NULL,
    approval_status SMALLINT CHECK(approval_status >= -1 AND approval_status <= 1)
);