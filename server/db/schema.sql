CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT uuid_generate_v4 () NOT NULL PRIMARY KEY,
    email TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
)

CREATE TABLE IF NOT EXISTS ledgers (
    id UUID DEFAULT uuid_generate_v4 () NOT NULL PRIMARY KEY,
    members UUID[] NOT NULL,
    transactions BIGINT[],
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
)

CREATE TABLE IF NOT EXISTS transations (
    id BIGINT NOT NULL PRIMARY KEY,
    debitor UUID NOT NULL REFERENCES users(id),
    creditor UUID NOT NULL REFERENCES users(id),
    amount FLOAT NOT NULL,
    note VARCHAR,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
)