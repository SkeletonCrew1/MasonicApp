# Masonic Application

## Overview
A distributed microservice application for tracking and managing supernatural sightings. Built with Docker and Docker Compose, the system provides role-based access, interactive map functionalities, and automated administrative tasks.

## Core Features
* **Interactive Map:** Submit and view sightings of supernatural creatures based on user ranks.
* **Role Management:** Rank system (Bronze, Silver, Gold, Inquisitor) with specific privileges.
* **Security & Auth:** JWT-based authentication, daily rotating passwords, and IP blacklisting.
* **Automated Services:** Scheduled tasks for password generation and Inquisitor selection.
* **Intrusion Response:** Silent redirects to a honeypot (`BirdWatching` app) for banned IPs.

## Service Architecture
The application utilizes a distributed architecture with specialized microservices:

* **`frontend`**: Vue 3 / Vite SPA. Handles the interactive map and user interfaces.
* **`auth-service`**: Go-based service managing user registration, JWT login, and IP ban verification.
* **`backend-django`**: Core administrative API (Python/Django). Handles IP banning, broadcasts, invites, and user promotions.
* **`posting`**: Manages CRUD operations for supernatural sightings and image uploads.
* **`mail-service`**: Asynchronous email delivery service for user notifications.
* **`daily-password`**: Generates daily access codes, stores them, and triggers mail notifications.
* **`daily-inquisitor`**: Randomly selects a Silver or Gold mason to be an Inquisitor and notifies them.
* **`birdwatching`**: A separate Flask application acting as a redirect honeypot for blacklisted users.

## Infrastructure & Databases
* **Databases (`main-db`, `map-db`, `bird-db`)**: PostgreSQL and MySQL instances isolating user data, coordinates, and honeypot records.
* **Migrations (`main-flyway`, `map-flyway`)**: Database schema management and version control using Flyway.

## Prerequisites
* Git
* Docker & Docker Compose
* Configured `.env` file in the root directory (ensure all required variables are set before running).

## Quick Start
1. Clone the repository:
   ```bash
   git clone [https://github.com/SkeletonCrew1/MasonicApp.git](https://github.com/SkeletonCrew1/MasonicApp.git)
   cd MasonicApp
2. Navigate to the root directory:
   ```bash
   cd MasonicApp
3. Build and start the containers:
   ```bash
   docker compose up --build
