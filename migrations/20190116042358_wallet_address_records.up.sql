CREATE TABLE wallet_address_records (
    uuid uuid NOT NULL DEFAULT gen_random_uuid(),
    actor TEXT NOT NULL REFERENCES actor_profile_records(actor) ON DELETE CASCADE,
    wallet_address TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (uuid)
);
CREATE INDEX wallet_address_records_index ON wallet_address_records (actor, wallet_address);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON wallet_address_records
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
