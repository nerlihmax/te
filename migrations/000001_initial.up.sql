CREATE TABLE IF NOT EXISTS accounts
(
    "id"   uuid DEFAULT gen_random_uuid() PRIMARY KEY UNIQUE,
    "name" text NOT NULL
);

CREATE INDEX IF NOT EXISTS accounts_id_index ON accounts (id);

CREATE TABLE IF NOT EXISTS entries
(
    "id"          serial PRIMARY KEY UNIQUE,
    "amount"      bigint NOT NULL CHECK ( amount > 0 ),
    "debit"       uuid   NOT NULL REFERENCES accounts (id) ON DELETE RESTRICT,
    "credit"      uuid   NOT NULL REFERENCES accounts (id) ON DELETE RESTRICT,
    "description" text,
    "timestamp"   timestamp DEFAULT now()
);

CREATE INDEX IF NOT EXISTS entries_id_index ON entries (id);