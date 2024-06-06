# RESTful API in Go with Clean Architecture To Be Growth Assessment

## Overview

This project implements a RESTful API using Go, following the principles of Clean Architecture. It includes PostgreSQL for the database, Redis for caching, JWT for authentication, and is containerized using Docker.

## Setup

1. **Clone the repository:**

```bash

git clone https://github.com/yourusername/yourproject.git

cd yourproject

```

2. \*\*Create a .env file with the following content:

```bash
DBUsername=go_assessment_user
DBPassword=KaCT3XW2UKgE
DBName=go_assessment_db
DBHost=localhost
DBPort=5432
RedisHost=redis
RedisPort=6379
JWTSecret=dxiyXWAcMnj66j3y4JPC7HhwUiW6CFqQecL7dpYU4qmbtj
```

3. **Build and run the Docker containers:**

```bash
docker-compose up --build
```

4. **Access the API at:**
   `http://localhost:8080`

## Endpoints

### Organizations

- **POST /api/organizations**: Create a new organization.
- **GET /api/organizations**: List all organizations.
- **GET /api/organizations/{id}**: Get an organization by ID.
- **PUT /api/organizations/{id}**: Update an organization.
- **DELETE /api/organizations/{id}**: Delete an organization.

### Users

- **POST /api/users**: Create a new user.
- **GET /api/users**: List all users.
- **GET /api/users/{id}**: Get a user by ID.
- **PUT /api/users/{id}**: Update a user.
- **DELETE /api/users/{id}**: Delete a user.

### Organization-User Association

- **POST /api/organizations/{orgId}/users/{userId}**: Add a user to an organization.
- **DELETE /api/organizations/{orgId}/users/{userId}**: Remove a user from an organization.

### Authentication

- **POST /api/login**: Obtain a JWT token by providing a valid email and password.

## Usage

### Authentication

To access protected routes, you need to include the JWT token in the `Authorization` header in the following format:
