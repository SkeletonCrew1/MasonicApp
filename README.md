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



* Check for all credentials in AWS Secret manager.
* Login in AWS acconut with `aws configure`.

## Create bucket for state files:
* Run `cd terraform/s3-bucket-for-state`
* Run `terraform init`
* Run `terraform plan`
* Run `terraform apply` ,check for any errors and approve if everything is okay

## Create bucket for photos:
* Run `cd ../s3-bucket-for-app`
* Run `terraform init`
* Run `terraform plan`
* Run `terraform apply` ,check for any errors and approve if everything is okay

## Create RDS:
* Run `cd ../rds`
* Run `terraform init`
* Run `terraform plan`
* Run `terraform apply` ,check for any errors and approve if everything is okay

## Change credentials:
- FLASK_VOTING_DATABASE_URL
- FLASK_MAIN_DATABASE_URL
- AUTH_URL
- PROMOTION_URL
- VOTING_DB_URL
- USERS_DB_URL
- VOTING_DB_PASSWORD
- MAP_DB_PASSWORD
- USERS_DB_PASSWORD

## Create EKS cluster:
* Run `cd ../eks-cluster`
* Run `terraform init`
* Run `terraform plan`
* Run `terraform apply` ,check for any errors and approve if everything is okay

## Start the application with helm:

* Make sure images are in ECR reposetories.If not,push images to ECR.

* Check if ECR repositories in MasonicApp/k8s/masonic-chart/masonic/values-stage.yaml

* Log into AWS EKS Cluster using (do not forget to change account id to actual value):
```
aws sts assume-role   --role-arn arn:aws:iam::<account id>:role/eks-admin   --role-session-name session
aws eks update-kubeconfig   --region eu-north-1   --name eks-cluster   --role-arn arn:aws:iam::<account id>:role/eks-admin
```
* Given you are inside project repostitory, run `cd k8s` and then run `helm install masonic ./masonic-chart/masonic/ --namespace application --create-namespace --values ./masonic-chart/masonic/values-stage.yaml` to start our whole application, If you make a deployment to production or dev stage change `./masonic-chart/masonic/values-stage.yaml` to the respective values file.

