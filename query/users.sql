
-- name: GetUserById :one
select * from mud.users as users where id = $1;

-- name: GetUserByName :one
select * from mud.users as users where name = $1;

-- name: CreateUser :one
insert into mud.users as users (name, password, level, being) values ($1, $2, $3, $4) returning *;

-- name: GetUserBeings :many
select * from mud.beings where owner = $1;

-- name: GetUserBeingNames :many
select name from mud.beings where owner = $1;