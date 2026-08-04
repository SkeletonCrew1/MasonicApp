# Django Backend Service | MasonicApp

## Overview
This directory contains the core backend service for the MasonicApp. Built with Django, this microservice exposes RESTful APIs to handle elevated privileges, user management, and administrative actions.

## Core Responsibilities
* **IP Management:** Manages the IP blacklist and verifies ban status for incoming requests.
* **Communications:** Handles broadcast messages.
* **Role Management:** Controls the promotion logic for user ranks (Bronze → Silver → Gold).
* **Access Control:** Processes and validates invitation links for new users.

## Directory Structure
* `api/`: The primary Django app containing the business logic (views), database schemas (models), and API route definitions.
* `core/`: The main Django project configuration directory (contains `settings.py`, root `urls.py`, etc.).
* `Dockerfile`: Containerization instructions for the service.
* `manage.py`: Django's command line utility for administrative tasks.
* `requirements.txt`: dependencies required for the service.

## Tech Stack
* **Framework:** Python / Django
* **Database Compatibility:** PostgreSQL (integrates with `main-db` and `map-db`)
* **Containerization:** Docker

## Running the Service

### Docker Compose
As part of the microservice architecture, this service is designed to be orchestrated via the root `docker-compose.yml`. 

Execute from the project root:
```bash
docker compose up backend-django --build