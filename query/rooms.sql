
-- name: GetRooms :many
select * from mud.rooms as rooms;

-- name: GetRoom :one
select * from mud.rooms where id = $1;

-- name: CreateUnlinkedRoom :one
insert into mud.rooms (name, description, objects) values ($1, $2, $3) returning *;

-- name: UpdateRoom :exec
update mud.rooms set name=$2, description=$3, img=$4, objects=$5, n=$6, s=$7, e=$8, w=$9, ne=$10, se=$11, sw=$12, nw=$13, u=$14, d=$15, i=$16, o=$17 where id=$1;