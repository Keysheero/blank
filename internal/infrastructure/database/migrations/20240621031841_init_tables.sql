-- +goose Up
-- +goose StatementBegin

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id            UUID PRIMARY KEY      DEFAULT gen_random_uuid(),
    email         VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at    TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE TABLE profiles
(
    user_id      SERIAL PRIMARY KEY,
    first_name   VARCHAR(50),
    last_name    VARCHAR(50),
    phone_number VARCHAR(15),
    address      TEXT
);

CREATE TABLE categories
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    description TEXT,
    parent_id   INT REFERENCES categories (id)
);

CREATE TABLE products
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100)   NOT NULL,
    description TEXT,
    price       DECIMAL(10, 2) NOT NULL,
    stock       INT            NOT NULL,
    category_id INT REFERENCES categories (id)
);

CREATE TABLE product_images
(
    id         SERIAL PRIMARY KEY,
    product_id INT REFERENCES products (id),
    image_url  TEXT NOT NULL
);

CREATE TABLE orders
(
    id         SERIAL PRIMARY KEY,
    user_id    UUID REFERENCES users (id),
    order_date TIMESTAMP   NOT NULL DEFAULT NOW(),
    status     VARCHAR(20) NOT NULL
);

CREATE TABLE order_items
(
    id         SERIAL PRIMARY KEY,
    order_id   INT REFERENCES orders (id),
    product_id INT REFERENCES products (id),
    quantity   INT            NOT NULL,
    price      DECIMAL(10, 2) NOT NULL
);

CREATE TABLE reviews
(
    id         SERIAL PRIMARY KEY,
    user_id    UUID REFERENCES users (id),
    product_id INT REFERENCES products (id),
    rating     INT       NOT NULL,
    comment    TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE cart
(
    id         SERIAL PRIMARY KEY,
    user_id    UUID REFERENCES users (id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE cart_items
(
    id         SERIAL PRIMARY KEY,
    cart_id    INT REFERENCES cart (id),
    product_id INT REFERENCES products (id),
    quantity   INT NOT NULL
);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cart_items;

DROP TABLE IF EXISTS cart;

DROP TABLE IF EXISTS reviews;

DROP TABLE IF EXISTS order_items;

DROP TABLE IF EXISTS orders;

DROP TABLE IF EXISTS product_images;

DROP TABLE IF EXISTS products;

DROP TABLE IF EXISTS categories;

DROP TABLE IF EXISTS profiles;

DROP TABLE IF EXISTS users;
-- +goose StatementEnd


