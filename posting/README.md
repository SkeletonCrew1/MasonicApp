# What is and how to run?
# What is?
This code is implementation of our posting service. Posting service allows us to create and get already existing sightings.
* `Dockerfile` is used to build our service in form of a container.
* `go.mod` and `god.sum` files are requirements that are used during building process.
* `internal/database/postgres.go` is a code we use to connect to our postgres database.
* `internal/handlers/sightings.go` is a code that is used to handle database requests used in this micro service.
* `internal/models/sighting.go` is an ORM we use for handling data regarding sigthings.
* `cmd/server/main.go` is our primary code and works as an entrypoint. This code builds a server that listens to requests and depending on the request performs required actions utilizing handlers.
## How to run?
* First download the repository.
* `cd` into repository folder
* Run `docker compose up --build` inside main repository folder (not posting)
## Prerequisites
* Have Git installed
* Have `Docker` and `Docker Compose` installed