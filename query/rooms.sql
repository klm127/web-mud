
-- name: GetRooms :many
select * from mud.rooms as rooms;

-- name: GetRoom :one
select * from mud.rooms where id = $1;

-- name: CreateUnlinkedRoom :one
insert into mud.rooms (name, description, objects) values ($1, $2, $3) returning *;