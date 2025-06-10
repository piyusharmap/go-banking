DROP TABLE IF EXISTS amount_transfer;

DO $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM pg_type WHERE typname = 'transfer_stage'
    ) THEN
        DROP TYPE transfer_stage;
    END IF;
END
$$;    