-- +goose Up 

create table users(
  id uuid PRIMARY KEY,
  created_at timestamp not null,
  updated_at timestamp not null,
  name text not null,
  unique(name)
);

-- +goose Down
drop table users;


