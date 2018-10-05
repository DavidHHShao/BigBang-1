package objective_config

const UPSERT_OBJECTIVE_COMMAND = `
INSERT INTO objectives 
(
  project_id,
  milestone_id,
  objective_id,
  content,
  total_rating,
  total_weight
)
VALUES 
(
  :project_id,
  :milestone_id,
  :objective_id,
  :content,
  :total_rating,
  :total_weight
)
ON CONFLICT (project_id, milestone_id, objective_id) 
DO
 UPDATE
    SET
        content = :content,
        total_rating = :total_rating,
        total_weight = :total_weight
    WHERE  objectives.objective_id = :objective_id and 
           objectives.project_id = :project_id and objectives.milestone_id = :milestone_id
RETURNING created_at;
`

const DELETE_OBJECTIVE_BY_IDS_COMMAND = `
DELETE FROM objectives
WHERE project_id = $1 and milestone_id = $2 and objective_id = $3;
`

const DELETE_OBJECTIVES_BY_PROJECT_ID_AND_MILESTONE_ID_COMMAND = `
DELETE FROM objectives
WHERE project_id = $1 and milestone_id = $2;
`

const QUERY_OBJECTIVE_BY_IDS_COMMAND = `
SELECT * FROM objectives
WHERE project_id = $1 and milestone_id = $2 and objective_id = $3;
`

const QUERY_OBJECTIVES_BY_PROJECT_ID_AND_MILESTONE_ID_COMMAND = `
SELECT * FROM objectives
WHERE project_id = $1 and milestone_id = $2
ORDER BY objective_id ASC;
`

const VERIFY_OBJECTIVE_EXISTING_COMMAND = `
select exists(select 1 from objectives where project_id = $1 and milestone_id = $2 and objective_id = $3);
`