CREATE TABLE IF NOT EXISTS customer(
    id SERIAL PRIMARY KEY,
    contact VARCHAR(15) UNIQUE NOT NULL,
    email VARCHAR(120) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS account(
    id SERIAL PRIMARY KEY,
    customer_id SERIAL NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50),
    account_number VARCHAR(20) UNIQUE NOT NULL,
    balance BIGINT DEFAULT 0 NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES customer(id) ON DELETE CASCADE
);

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transfer_stage') THEN
        CREATE TYPE transfer_stage AS ENUM ('PENDING', 'SUCCESS', 'FAILED');
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS amount_transfer(
    id SERIAL PRIMARY KEY,
    sender_account_id INT NOT NULL, 
    receiver_account_id INT NOT NULL,
    amount BIGINT NOT NULL CHECK (amount > 0),
    stage transfer_stage NOT NULL DEFAULT 'PENDING',
    remark VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sender_account_id) REFERENCES account(id) ON DELETE CASCADE,
    FOREIGN KEY (receiver_account_id) REFERENCES account(id) ON DELETE CASCADE,
);