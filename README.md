# API & CLI in Golang:

This repository contains a Golang API with a PostgreSQL database and a CLI to interact with the API, organized as follows:

- `app/`: Contains the Go application code, its Dockerfile, and the environment variables file.
- `postgres-setup/`: Contains the SQL script for configuring the PostgreSQL database and its Dockerfile.
- `cli/`: Contains the CLI source code, its Dockerfile, and the necessary files for its operation.
- `docker-compose.yml`: Docker Compose configuration file for building and running the entire application.
- `.env`: Environment variables file to store the database connection parameters.

## Prerequisites

- Docker
- Docker Compose

## How to build and run the API and CLI

1. Clone the repository:

```bash
git clone https://gitlab.com/your-username/golang_project_ilies_sylvain.git
cd golang_project_ilies_sylvain
```

2. Build and run the application using Docker Compose:

```bash
docker-compose up -d --build
```

This command will build the Go application, PostgreSQL, and CLI containers, then run them together. Your application will be accessible at http://localhost:8080.

To stop the containers, press Ctrl+C or run:

```bash
docker-compose down -v
```
## API Endpoints

* `/users`: Manage users (GET, POST, PUT, DELETE).
* `/roles`: Manage roles for users (GET, POST, PUT, DELETE).
* `/groups`: Manage user groups (GET, POST, PUT, DELETE).
* `/auth`: Manage user authentication using JWT (POST).

### /users

* `GET /users`: Retrieve the list of users.
* `POST /users`: Create a new user.
* `PUT /users/:id`: Update an existing user with the specified ID.
* `DELETE /users/:id`: Delete a user with the specified ID.

### /roles

* `GET /roles`: Retrieve the list of roles.
* `POST /roles`: Create a new role.
* `PUT /roles/:id`: Update an existing role with the specified ID.
* `DELETE /roles/:id`: Delete a role with the specified ID.

### /groups

* `GET /groups`: Retrieve the list of groups.
* `POST /groups`: Create a new group.
* `PUT /groups/:id`: Update an existing group with the specified ID.
* `DELETE /groups/:id`: Delete a group with the specified ID.

### /auth

* `POST /auth`: Authenticate a user and return a JWT token.
* `POST /signup`: Create a user in the DB with email + password.
* `POST /login`: Authenticate a user + return the JWT token.
* `POST /validate`: Retrieve the JWT token for analysis and securing access routes.

## Using the CLI

To use the CLI, it is strongly recommended to create an alias:

```bash
alias cli="docker exec -it cli ./cli"
```

To be able to run commands inside the container:

```bash
cli your-command [args]
```

Otherwise, you can run commands inside the container using docker exec:

```bash
docker exec -it cli ./cli your-command [args]
```

Replace your-command and [args] with the appropriate command and arguments for your CLI application.

### Available commands

* `login`: Log in as a user and retrieve an authentication JWT token and a refresh token.
    * Flags:
        * `--email`: User's email address.
        * `---password`: User's password.
* `refresh`: Refresh an authentication JWT token using a refresh token.
    * Flags:
        * `--refresh_token`: The refresh token.
* `logout`: Log out and delete an authentication JWT token and a refresh token.
    * Flags:
        * `--access_token`: The authentication JWT token.
* `users list`: List all users.
* `users get [user_id]`: Retrieve a specific user.
* `users create`: Create a new user.
    * Flags:
        * `--email`: User's email address.
        * `--password`: User's password.
        * `--name`: User's full name.
* `users update [user_id]`: Update an existing user.
    * Flags:
        * `--email`: User's new email address.
        * `--password`: User's new password.
        * `--name`: User's new full name.


