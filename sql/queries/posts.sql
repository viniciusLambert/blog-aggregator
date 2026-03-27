-- name: CreatePost :one
insert into posts (id, created_at, updated_at, title, url, description, published_at, feed_id )
values (
  $1, 
  $2, 
  $3, 
  $4, 
  $5, 
  $6,
  $7,
  $8
)
RETURNING *;


-- name: GetPostForUser :many
select *
from posts p
inner join feeds f
on p.feed_id = f.id
where f.user_id = $1 
order by p.published_at desc
limit($2);
