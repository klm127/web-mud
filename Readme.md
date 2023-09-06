# Overview

This is a web-based MUD. A MUD is a Multi-User-Dungeon. They were text based multiplayer RPGs that were a precursor to modern MMOs. They were typically played over Telnet. This one will simulate some kind of terminal with a limited GUI in the user's browser.

I set it up with the Gostgres stack. That is, Go for the HTTP server and Postgres for the Database.

It will use TypeScript to generate JS needed for front end stuff.

It should use [WebSockets](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API) to communicate with a user once they are playing.

These modern languages are strongly typed and should help manage what could become a complex project.

# Setting Up

You need to install [docker](https://www.docker.com/) and [make](https://community.chocolatey.org/packages/make)

You can install make for windows in an elevated powershell using `choco install make`. This utilizes the [chocolatey](https://chocolatey.org/) package manager.

You need to install [Go](https://go.dev/).

# Overview

- A Postgres database will store game and user information. This database runs in a docker container.
- A Go process, which will eventually be put in its own container, connects to the database. That process creates an HTTP Server that listens for HTTP Requests.
- As HTTP Requests are received, the server servers pages rendered from the files in `server/templates` which use assets in `server/static`.

# Make commands

Run these make commands to get started.

`make db` - Create the mysql database
`make go.dev` - Run the go program with the same env vars used to create the database

_Optional_:

`make db.admin.run` - Create the pgAdmin server

# Developing - Front End

If you want to do front end dev, you will probably want to create javascript. This project uses TypeScript. 

If you run `make ts.watch` the typescript transpiler will start in edit mode. As you create and edit typescript files, they will be modified and created in `server/static/js`.

You will only need to link the base js file that imports whatever you need - JS module imports will take care of the rest.

Example:

*the typecript*
```ts
// typescript/myfunc.ts
export function hello() {
    return "Hi!"
}

// typescript/myapp.ts
import {hello} from "./myfunc.js"

const html_element = document.createElement("div")
div.textContent = hello()
document.body.append(div)

```

*in the template*
```html
<link rel="/js/myapp.js" type="text/javascript">
```

# Developing - Back End

You can use pgAdmin to edit tables in the database. Use the `make db.admin.run` command to start it.

Once it starts, connect to the database using the credentials in `dev.env`. From here you can use the pgAdmin GUI to edit the database.

Once you have made some changes, export the database as a .sql file. Place or replace a file in `.docker/postgres`. This is where the setup scripts for the database are stored and are what will be run next time you create the container. You'll need to do this to propogate your changes to other devs.

You can run `make db.rb` to rebuild the database and ensure you configured the exported .sql file; you might have to make some small changes.

To access the database in Go, a program called [sqlc](https://pkg.go.dev/github.com/kyleconroy/sqlc) is used to automagically generate Go code. Write a query in /query. Based on the comment above the query, code will be generated in `db/dbg`

To use the code generated from anywhere in the program, simply use the global Store struct in `db`

Example: 
```go
import (
	"github.com/pwsdc/web-mud/db"
)
// ...
rooms := db.Store.Query.GetRooms(context.Background())
```

You will probably only ever need to pass an empty `context.Background()` as the first parameter; I don't suspect context functionalities will be much needed in database calls.

To add an endpoint URL, you'll need to create a function somewhere sensible in `server`. That function will have to be connected to the router. See [gin-gonic](https://pkg.go.dev/github.com/gin-gonic/gin) for more information. 

