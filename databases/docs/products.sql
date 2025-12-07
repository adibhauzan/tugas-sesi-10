CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    type_id INT NOT NULL REFERENCES product_types(id) ON UPDATE CASCADE,
    price NUMERIC(12,2) NOT NULL,
    stock INT NOT NULL,

    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
