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
RETURNING created_at;
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
SELECT COALESCE(update_count, 0) FROM posts
WHERE post_hash = $1;
`

const VERIFY_POSTHASH_EXISTING_COMMAND = `
select exists(select 1 from posts where post_hash =$1);
`

const QUERY_POST_TYPE_COMMAND = `
SELECT post_type FROM posts
WHERE post_hash = $1;
`
