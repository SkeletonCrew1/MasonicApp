# What is and how to run?
# What is?
This code is implementation of our daily inquisitor service. This code allows us to select and assign a random mason as an inquisitor every day.
* `Dockerfile` is used to build our service in form of a container.
* `go.mod` and `god.sum` files are requirements that are used during building process.
* `internal/database/postgres.go` is a code we use to connect to our postgres database.
* `internal/handlers/code.go` is a code that is used to handle database code writing requests used in this micro service.
* `internal/handlers/users.go` is a code that is used to handle database users related requests used in this micro service. Used to notify mail service who should get the new password.
* `internal/models/user.go` is an ORM we use for handling data regarding users.
* `cmd/password/main.go` is our primary code and works as an entrypoint. This code builds a server that selects a random golden or silver user as an inquisitor every 24 hours.
## How to run?
* First download the repository.
* `cd` into repository folder
* Run `docker compose up --build`
## Prerequisites
* Have Git installed
* Have `Docker` and `Docker Compose` installed