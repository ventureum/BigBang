package post_config

const UPSERT_POST_COMMAND = `
INSERT INTO posts 
(
  actor, 
  board_id, 
  parent_hash, 
  post_hash, 
  post_type, 
  content
)
VALUES 
(
  :actor, 
  :board_id, 
  :parent_hash, 
  :post_hash, 
  :post_type, 
  :content
)
ON CONFLICT (post_hash) 
DO
 UPDATE
    SET actor = :actor,
        board_id = :board_id,
        parent_hash = :parent_hash,
        post_type = :post_type,
        content = :content,
        update_count = posts.update_count + 1
    WHERE posts.post_hash = :post_hash
RETURNING updated_at;
`

const DELETE_POST_COMMAND = `
DELETE FROM posts
WHERE post_hash = $1;
`

const QUERY_POST_COMMAND = `
SELECT * FROM posts
WHERE post_hash = $1;
`

const QUERY_POST_UPDATE_COUNT_COMMAND = `
SELECT update_count FROM posts
WHERE post_hash = $1;
`
