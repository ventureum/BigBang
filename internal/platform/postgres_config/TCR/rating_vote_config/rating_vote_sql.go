package rating_vote_config

const UPSERT_RATING_VOTE_COMMAND = `
INSERT INTO rating_votes 
(
  project_id,
  milestone_id,
  objective_id,
  voter,
  rating,
  weight
)
VALUES 
(
  :project_id,
  :milestone_id,
  :objective_id,
  :voter,
  :rating,
  :weight
)
ON CONFLICT (project_id, milestone_id, objective_id, voter) 
DO
 UPDATE
    SET
        rating = :rating,
        weight = :weight
    WHERE rating_votes.project_id = :project_id and 
          rating_votes.milestone_id = :milestone_id and
          rating_votes.objective_id = :objective_id and
          rating_votes.voter = :voter
RETURNING created_at;
`

const DELETE_RATING_VOTE_BY_IDS_AND_VOTER_COMMAND = `
DELETE FROM rating_votes
WHERE project_id = $1 and milestone_id = $2 and objective_id = $3 and voter = $4;
`

const DELETE_RATING_VOTE_BY_IDS_COMMAND = `
DELETE FROM rating_votes
WHERE project_id = $1 and milestone_id = $2 and objective_id = $3;
`

const QUERY_RATING_VOTE_BY_IDS_AND_VOTER_COMMAND = `
SELECT * FROM rating_votes
WHERE project_id = $1 and milestone_id = $2 and objective_id = $3 and voter = $4;
`

const QUERY_RATING_VOTES_BY_IDS_COMMAND = `
SELECT * FROM rating_votes
WHERE project_id = $1 and milestone_id = $2 and objective_id = $3;
`

const VERIFY_RATING_VOTE_EXISTING_COMMAND = `
select exists(select 1 from rating_votes where project_id = $1 and milestone_id = $2 and objective_id = $3 and voter = $4);
`

const PAGINATION_QUERY_RATING_VOTE_LIST_COMMAND = `
SELECT * FROM rating_votes
WHERE id <= $4 and project_id = $1 and milestone_id = $2 and objective_id = $3
ORDER BY id DESC
LIMIT $5;
`

const QUERY_RATING_VOTE_LIST_COMMAND = `
SELECT * FROM rating_votes
WHERE project_id = $1 and milestone_id = $2 and objective_id = $3
ORDER BY id DESC
LIMIT $4;
`
