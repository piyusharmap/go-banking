DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_type WHERE typname = 'transfer_stage'
    ) THEN
        CREATE TYPE transfer_stage AS ENUM ('PENDING', 'COMPLETED', 'FAILED');
    END IF;
END
$$;

CREATE TABLE IF NOT EXISTS amount_transfer(
    id SERIAL PRIMARY KEY,
    sender_account_id INT NOT NULL,
    receiver_account_id INT NOT NULL,
    amount BIGINT NOT NULL CHECK (amount > 0),
    stage transfer_stage NOT NULL,
    remark VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sender_account_id) REFERENCES account(id) ON DELETE CASCADE,
    FOREIGN KEY (receiver_account_id) REFERENCES account(id) ON DELETE CASCADE
);