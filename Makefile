# ------------------------------------------------------------
# 		Environment variables
# ------------------------------------------------------------

 # if this is set to 'production', production.env file will be included.
 # otherwise, dev.env fill be included
 
mode=development

 # either .env will likely overwrite the following defaults:

# ------------------------------------------------------------
# 				Environment Defaults
# ------------------------------------------------------------

 # database name
db=sdcmud
 # database super user
db_user=postgres
 # database password
db_pw=bwgpw
 # database image
db_img=postgres:15.4-bullseye
 # database container
db_container=sdcmuddb
 # port to be exposed on db
db_port=5432
# port to be exposed for http server
http_port=80

# ------------------------------------------------------------
# 				Non-Environment Vars
# ------------------------------------------------------------
 # image to use for pgadmin
pgadmin_img=dpage/pgadmin4:7.5
 # container name for pgadmin
pgadmin_container=sdcmudadmin
 # port for pgadmin
pgadmin_port=3000

docker_net_name=bwgnet


ifeq ($(mode), production)
	include production.env
else 
	include dev.env
endif

# ------------------------------------------------------------
# 		Combination commands
# ------------------------------------------------------------
 # make all
all: network.create db db.admin.run sqlc 
 # remove all
remove: db.rm db.admin.rm network.rm
 # rebuild all except network
rb: db.rb db.admin.rb sqlc
 # start commands
start: db.start db.admin.start

# ------------------------------------------------------------
# 		Postgres Database Image and Container Build Commands
# ------------------------------------------------------------

 # Build the image, then run it.
db: db.run db.conn

 # 'database run' runs the mysql docker container from the image that was built.
db.run:
	docker run -p 5432:$(db_port) \
	-v '$(CURDIR)/.docker/postgres:/docker-entrypoint-initdb.d' \
	-e POSTGRES_USER=$(db_user) \
	-e POSTGRES_PASSWORD=$(db_pw) \
	-e POSTGRES_DATABASE=$(db) \
	--name $(db_container) -d $(db_img)

 # 'database remove' removes the mysql docker container.
db.rm:
	docker rm $(db_container) --force

 # 'database rebuild' removes the mysql docker container, then rebuilds the image, then runs the container again with the new image.
db.rb: db.rm db

 # 'databash bash' Runs /bin/bash on the database container, allowing access to the database filesystem. Useful for debugging.
db.bash:
	docker exec -it $(db_container) /bin/bash

 # 'database pq' Runs pq on the database container, allowing access to the interactive sql repl. Useful for debugging. Some useful commands: "show tables;", "show databases;" "use betacusc_typ_livedb;" "select * from tablename;" etc.
db.pq:
	docker exec -it ${db_container} psql -U ${db_user} 

 # 'database image help' Prints help to the command line about the mysql image. You could > this into a file to make it more readable. 
db.image.help:
	docker run -it --rm ${db_image} --verbose --help

db.conn:
	docker network connect $(docker_net_name) $(db_container)

db.dc:
	docker network disconnect $(docker_net_name) $(db_container)

db.dump:
	docker exec -it $(db_container) mkdir -p /tmp/backup
	docker exec -it $(db_container) \
	pg_dump --file "/tmp/backup/full.sql" --format=p --section=pre-data --section=data --section=post-data --exclude-schema public --schema mud --inserts --on-conflict-do-nothing --create --clean --if-exists --verbose \
	-U $(db_user) -d $(db) -n public
	docker cp $(db_container):/tmp/backup/full.sql "$(CURDIR)/.docker/postgres/bup.sql"

# ------------------------------------------------------------
# 		pgAdmin Helper Image and Container Build Commands
# ------------------------------------------------------------

db.admin: db.admin.run db.admin.conn

 # run pgAdmin container
db.admin.run:
	docker run \
	--name $(pgadmin_container) \
	-e PGADMIN_DEFAULT_EMAIL=admin@bwg.net \
	-e PGADMIN_DEFAULT_PASSWORD=password \
	-e MASTER_PASSWORD_REQUIRED=false \
	-e PGADMIN_LISTEN_PORT=80 \
	-p 3000:80 \
	-d $(pgadmin_img)

 # remove phpmyadmin container
db.admin.rm:
	docker rm $(pgadmin_container) --force

 # rebuild phpmyadmin
db.admin.rb: db.admin.rm db.admin.run db.admin.conn

 # connect to the docker network
db.admin.conn:
	docker network connect $(docker_net_name) $(pgadmin_container)

 # disconnect from the docker network
db.admin.dc:
	docker network disconnect $(docker_net_name) $(pgadmin_container)

 # 'pgadmin bash' Runs /bin/bash on the pgadmin container, allowing access to the database filesystem. Useful for debugging.
db.admin.bash:
	docker exec -it $(pgadmin_container) /bin/sh

# ------------------------------------------------------------
# 		Code Generation and Migration
# ------------------------------------------------------------

 # Runs sqlc to generate ORM code from database schema and queries
sqlc:
	docker run --rm -v "$(CURDIR):/src" -w /src kjconroy/sqlc generate

# ------------------------------------------------------------
# 		Network commands
# ------------------------------------------------------------
network.create:
	docker network create -d bridge $(docker_net_name)

network.rm:
	docker network rm $(docker_net_name)

network.dcall: db.admin.dc db.dc 

network.conn: db.admin.conn db.conn

network.rb: network.dcall network.rm network.create network.conn

network.insp: 
	docker network inspect $(docker_net_name)

# ------------------------------------------------------------
# 		Go run commands
# ------------------------------------------------------------

 # run go using dev environment variables
go.dev: 
	go run . dev.env

 # run go using production environment variables
go.prod:
	go run . production.env

 # build executable
go.build:
	go build .

# ------------------------------------------------------------
# 		Typescript/JS commands
# ------------------------------------------------------------

ts.watch:
	cd typescript & npx tsc --watch

ts.build.clean: typescript/node_modules
	del /q server\static\js
	cd typescript & del /q dist & npx tsc 

ts.build: typescript/node_modules
	cd typescript & npx tsc

typescript/node_modules:
	cd typescript & yarn install
