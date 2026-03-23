-- +goose Up
create table feeds(
  id uuid PRIMARY KEY,
  created_at timestamp not null,
  updated_at timestamp not null,
  name text not null,
  url text not null,
  user_id uuid not null references users(id) on delete cascade,
  unique(url)
);

-- +goose Down
drop table feed;
