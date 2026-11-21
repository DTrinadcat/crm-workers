-- +goose Up
create table if not exists workers (
  id uuid primary key,
  name text not null,
  surname text not null,
  role text not null,
  office text not null,
  number text not null,
  additional_number text not null default '',
  email text not null unique,
  address text not null,
  is_admin boolean not null default false,
  salary integer not null default 0
);

-- +goose Down
drop table if exists workers;