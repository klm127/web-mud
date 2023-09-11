
-- name: GetUserById :one
select * from mud.users as users where id = $1;

-- name: GetUserByName :one
select * from mud.users as users where name = $1;

-- name: CreateUser :one
insert into mud.users as users (name, password, level) values ($1, $2, $3) returning *;