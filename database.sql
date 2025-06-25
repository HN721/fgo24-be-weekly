-- Active: 1750554601825@@127.0.0.1@5434@erd
CREATE DATABASE erd;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    images VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE balance (
    id SERIAL PRIMARY KEY,
    saldo int,
    last_updated TIMESTAMP,
    id_user int REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE pin (
    id SERIAL PRIMARY KEY,
    pin int,
    id_user int REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE topup (
    id SERIAL PRIMARY KEY,
    amount int,
    date_transactions DATE,
    id_user int REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE transaction (
    id SERIAL PRIMARY KEY,
    amount int,
    status VARCHAR(20) CHECK (
        status in (
            'pending',
            'completed',
            'failed'
        )
    ),
    date_transaction DATE,
    method VARCHAR(255),
    id_user int REFERENCES users (id) ON DELETE CASCADE
);