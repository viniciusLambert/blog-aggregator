-- +goose Up
create table posts (
  id uuid PRIMARY KEY,
  created_at timestamp not null,
  updated_at timestamp not null,
  title text not null,
  url text not null,
  description text,
  published_at timestamp,
  feed_id uuid not null references feeds(id) on delete cascade,
  unique(url)
);
-- +goose Down
drop table posts;
