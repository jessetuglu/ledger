CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT uuid_generate_v4 () NOT NULL PRIMARY KEY,
    email TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE(email)
);

CREATE TABLE IF NOT EXISTS ledgers (
    id UUID DEFAULT uuid_generate_v4 () NOT NULL PRIMARY KEY,
    title VARCHAR NOT NULL,
    members UUID[] NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE(title)
);

CREATE TABLE IF NOT EXISTS transactions (
    id BIGINT NOT NULL PRIMARY KEY,
    ledger UUID NOT NULL REFERENCES ledgers(id),
    debitor UUID NOT NULL REFERENCES users(id),
    creditor UUID NOT NULL REFERENCES users(id),
    date TIMESTAMP NOT NULL,
    amount FLOAT NOT NULL,
    note VARCHAR,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);