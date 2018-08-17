package purchase_reputations_record_config


const TABLE_SCHEMA_FOR_PURCHASE_REPUTATIONS_RECORDS = `
CREATE TABLE purchase_reputations_records (
    uuid uuid NOT NULL DEFAULT gen_random_uuid(),
    purchaser TEXT NOT NULL,
    vetx BIGINT NOT NULL,
    reputations BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (uuid)
);
`

const TABLE_NAME_FOR_PURCHASE_REPUTATIONS_RECORDS = "purchase_reputations_records"
