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

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON redeem_block_info_records
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

INSERT INTO redeem_block_info_records
(
  redeem_block,
  executed_at
)
VALUES
( (CAST (EXTRACT(EPOCH FROM NOW()) AS BIGINT)) / (60 * 60 * 24 * 7) + 1,
  to_timestamp( ((CAST (EXTRACT(EPOCH FROM NOW()) AS BIGINT)) / (60 * 60 * 24 * 7) + 1) * (60 * 60 * 24 * 7)))
ON CONFLICT (redeem_block)
   DO NOTHING;