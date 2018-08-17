package session_record_config

const UPSERT_SESSION_RECORD_COMMAND = `
INSERT INTO session_records
(
  actor,
  post_hash,
  start_time, 
  end_time, 
  content
)
VALUES 
(
  :actor,
  :post_hash, 
  :start_time, 
  :end_time, 
  :content
)
ON CONFLICT (post_hash) 
DO
 UPDATE
    SET actor = :actor,
        start_time = :start_time,
        end_time = :end_time,
        content = :content
    WHERE session_records.post_hash = :post_hash
RETURNING updated_at;
`

const DELETE_SESSION_RECORD_COMMAND = `
DELETE FROM session_records
WHERE post_hash = $1;
`

const QUERY_SESSION_RECORD_COMMAND = `
SELECT * FROM session_records
WHERE post_hash = $1;
`
