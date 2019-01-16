CREATE TABLE purchase_mps_records (
    uuid uuid NOT NULL DEFAULT gen_random_uuid(),
    purchaser TEXT NOT NULL REFERENCES actor_profile_records(actor),
    vetx BIGINT NOT NULL,
    mps BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (uuid)
);
