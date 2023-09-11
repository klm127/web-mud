
-- name: GetUserById :one
select * from users where id = $1;

-- name: GetUserByName :one
select * from users where name = $1;

-- name: CreateUser :one
insert into users (name, password, level) values ($1, $2, $3) returning *;