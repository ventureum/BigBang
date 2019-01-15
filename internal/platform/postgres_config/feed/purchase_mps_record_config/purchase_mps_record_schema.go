package purchase_mps_record_config

const TABLE_SCHEMA_FOR_PURCHASE_MPS_RECORDS = `
CREATE TABLE purchase_mps_records (
    uuid uuid NOT NULL DEFAULT gen_random_uuid(),
    purchaser TEXT NOT NULL REFERENCES actor_profile_records(actor),
    vetx BIGINT NOT NULL,
    mps BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (uuid)
);
`

const TABLE_NAME_FOR_PURCHASE_MPS_RECORDS = "purchase_mps_records"
