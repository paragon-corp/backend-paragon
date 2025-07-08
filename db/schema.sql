-- db/schema.sql
CREATE TABLE coffees (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    price NUMERIC NOT NULL
);

CREATE TABLE orders (
    id TEXT PRIMARY KEY,
    coffee_id TEXT NOT NULL REFERENCES coffees(id),
    status TEXT NOT NULL
);
