-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
  $1, 
  $2, 
  $3, 
  $4, 
  $5, 
  $6
)
RETURNING *;

-- name: FetchFeeds :many
select *
from feeds;

-- name: FetchFeedsWithUserName :many

select f.*, u.name as user_name
from feeds f
left join users u
on f.user_id = u.id;

-- name: GetFeedsByUrl :one
select *
from feeds
where url = $1;


-- name: ClearFeeds :exec
delete from feeds;
