package post_replies_record_config

const UPSERT_POST_REPLIES_RECORD_COMMAND = `
INSERT INTO post_replies_records
(
  post_hash,
  reply_hash
)
VALUES 
(
  :post_hash, 
  :reply_hash
)
ON CONFLICT (post_hash, reply_hash) 
DO NOTHING;
`

const DELETE_POST_REPLIES_RECORD_COMMAND = `
DELETE FROM post_replies_records
WHERE post_hash = $1;
`

const QUERY_POST_REPLIES_COMMAND = `
SELECT reply_hash FROM post_replies_records
WHERE post_hash = $1;
`

const QUERY_POST_REPLIES_COUNT_COMMAND = `
SELECT COUNT(*) FROM post_replies_records
WHERE post_hash = $1;
`
