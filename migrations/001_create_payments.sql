CREATE TABLE payments (
    id         UUID            PRIMARY KEY,
    order_id   UUID            NOT NULL,
    amount     NUMERIC (12, 2) NOT NULL,
    status     VARCHAR (20)    NOT NULL,
    created_at TIMESTAMP       NOT NULL,
    updated_at TIMESTAMP       NOT NULL
);