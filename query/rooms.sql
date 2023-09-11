
-- name: GetRooms :many
select * from mud.rooms as rooms;

-- name: GetRoom :one
select * from mud.rooms where id = $1;
