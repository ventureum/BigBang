CREATE TABLE redeem_block_info_records (
    redeem_block BIGINT NOT NULL DEFAULT 0,
    total_enrolled_milestone_points BIGINT NOT NULL DEFAULT 0,
    token_pool BIGINT NOT NULL DEFAULT 10000,
    executed_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (redeem_block)
);
CREATE INDEX redeem_block_info_records_index ON redeem_block_info_records (redeem_block, total_enrolled_milestone_points, token_pool);
