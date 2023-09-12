
-- name: GetBeingByName :one
select * from mud.beings where name = $1;

-- name: GetBeingById :one
select * from mud.beings where id = $1;

-- name: CreateBeing :one
insert into mud.beings (name, description, room) values ($1, $2, $3) returning *;

-- name: UpdateBeingOwner :exec
update mud.beings set owner=$2 where id=$1;

-- name: UpdateBeing :exec
update mud.beings set description=$2, room=$3 where id=$1;

-- name: DeleteBeing :exec
delete from mud.beings where id = $1;