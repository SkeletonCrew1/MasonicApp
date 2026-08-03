# Masonic Application
## Application Overview
This repository is our implementation of the masonic application running inside Docker containers with the help of Docker Compose.
Our application features:
* Interactive map where users (depending on the rank) can submit sightings of supernatural creatures or events.
* User ranking.
* Registration, login and authorization.
* Daily password generation.

## Breakdown of folders structure
Our application works in a distributed fashion, we have seperate frontend, backend, databases and a lot of microservices.

Here is a breakdown of services:
* `frontend` is our map and frontend side of the application.
* `auth-service` is our service that handles user registration, login and authorization.
* `posting` is our service that handles everything related to sightings, namely: sighting submission, sighting fetching whether all sightings or a specific one.
* `mail-service` is our service that accepts incoming users and body and then sends the text in body to the users.
* `daily-password` is a service that handles generation of new daily password, inserting it into our database and notifying mail about users who must be notified.
* `daily-inquisitor` is a service that selects a random silver or golden mason to be an inquisitor, changes their status to inquisitor in a database and makes sure that this user is notified.
* `db` is a folder that contains all databases configuration scripts.
* `backend-django` is a service that handles big part of our functionality, such as: sending invites, banning users by their IP, sending broadcast messages and promoting users.
* `docker-compose.yml` is our main entrypoint of our application. Running this allows to start the whole application.


## How to run and prerequisites
To run this application you will need to have `Git`, `Docker` and `Docker Compose` installed.
To run follow the next steps:
* Run `git clone https://github.com/SkeletonCrew1/MasonicApp.git` and wait for the repository to be downloaded.
* Run `cd MasonicApp`.
* Run `docker compose up --build`.