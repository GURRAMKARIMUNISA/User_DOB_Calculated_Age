Go User API
A simple RESTful API for managing user data (Create, Read, Update, Delete) built with Go (Fiber framework), PostgreSQL, and `sqlc` for type-safe database interactions, all containerized with Docker Compose.

## Features

* **User Management:** Create, retrieve, update, and delete user records.
* **Database:** PostgreSQL for persistent data storage.
* **ORM/SQL Generation:** `sqlc` for generating type-safe Go code from SQL queries.
* **Web Framework:** Fiber for building fast and expressive HTTP APIs.
* **Logging:** Structured logging with `go.uber.org/zap`.
* **Configuration:** Environment-based configuration loading.
* **Containerization:** Docker Compose for easy setup and deployment of the application and database.
* **User Age Calculation:** Automatically calculates age based on DOB.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

* **Git:** For cloning the repository.
* **Docker Desktop:** Includes Docker Engine and Docker Compose. (Recommended for Windows/macOS)
    * [Download Docker Desktop](https://www.docker.com/products/docker-desktop)
* **Go:** Version 1.24 or higher.
    * [Download Go](https://golang.org/dl/)
* **`sqlc`:** Go SQL compiler.
    * Install with: `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`
* **`psql` client:** PostgreSQL command-line client (usually comes with PostgreSQL installation, or can be installed separately).
* **`curl` or Postman/Insomnia:** For testing API endpoints.

## Getting Started

Follow these steps to get the API up and running on your local machine.

### 1. Clone the Repository


git clone https://github.com/GurramKarimunisa/go-user-api.git
cd go-user-api

**(Replace `GurramKarimunisa` with your actual GitHub username)**

### 2. Set Up Environment Variables

Create a `.env` file in the root of the project (`go-user-api/`) with the following content:


DATABASE_URL=postgres://user_api_user:user_api_password@db:5432/user_api_db?sslmode=disable
PORT=3000
ENVIRONMENT=development

* **Note:** We are using `user_api_user` and `user_api_password` as defined in `docker-compose.yml`.
* `db` is the hostname of the PostgreSQL service within the Docker Compose network.

### 3. Initialize Go Modules and Generate `sqlc` Code

Run these commands from the project root (`go-user-api/`):


go mod tidy
sqlc generate


### 4. Start Docker Compose Services

This command will build your Go application Docker image, start the PostgreSQL database, and then start your Go API service. It will wait for the database to be ready before starting the API.


docker-compose up --build -d

* `--build`: Rebuilds the Go application image if changes are detected.
* `-d`: Runs the containers in detached mode (in the background).

### 5. Apply Database Migrations

Once the Docker services are running, you need to create the `users` table in your PostgreSQL database.

Open a **new Command Prompt (CMD) as Administrator** and run the following command. When prompted for the password, enter `user_api_password`.


set PGPASSWORD=user_api_password
psql -h 127.0.0.1 -p 5432 -U user_api_user -d user_api_db -f db/migrations/001_create_users_table.sql

You should see `CREATE TABLE` as the output, indicating success.

### 6. Verify Services (Optional)

Check if both services are running:

docker-compose ps

You should see `db` and `app` services listed as `Up` or `healthy`.

To view the application logs:

docker-compose logs app


## API Endpoints

The API will be accessible at `http://localhost:3000`. Use `curl` (remember to escape double quotes for Windows CMD) or a tool like Postman/Insomnia to test the endpoints.

**Note:** For `curl` on Windows CMD, use double quotes for the data payload and escape inner double quotes with a backslash (`\"`).

### 1. Create a New User (POST)


curl -X POST -H "Content-Type: application/json" -d "{\"name\": \"Alice\", \"dob\": \"1990-05-10\"}" http://localhost:3000/users

* **Response (Success):** `{"id":1,"name":"Alice","dob":"1990-05-10T00:00:00Z","age":34}` (ID and age may vary)

### 2. Get All Users (GET)


curl http://localhost:3000/users


### 3. Get User by ID (GET)


curl http://localhost:3000/users/1

(Replace `1` with the actual user ID)

### 4. Update User (PUT)


curl -X PUT -H "Content-Type: application/json" -d "{\"name\": \"Alice Updated\", \"dob\": \"1991-03-15\"}" http://localhost:3000/users/1

(Replace `1` with the actual user ID)

### 5. Delete User (DELETE)


curl -X DELETE http://localhost:3000/users/1

(Replace `1` with the actual user ID)

## Clean Up

To stop and remove all Docker containers and the associated database data volume:


docker-compose down --volumes


## Troubleshooting

* **`password authentication failed for user`**:
    * Ensure your `docker-compose.yml` has `POSTGRES_PASSWORD: user_api_password` for the `db` service.
    * Make sure you are typing `user_api_password` correctly at the `psql` prompt (characters won't show).
    * **Most likely issue:** You have another PostgreSQL server running on your host machine on port 5432. Check Task Manager (`postgres.exe` or `pg_ctl.exe`) and Windows Services (`services.msc`) to stop and disable it. Then run `docker-compose down --volumes` and `docker-compose up --build -d` again.
    * Consider setting `PGPASSWORD=user_api_password` in your CMD session before running `psql` to bypass interactive typing.
* **`connection refused`**:
    * The database might not be fully ready. The `healthcheck` in `docker-compose.yml` should prevent this, but allow some extra time.
    * Ensure your `app` service's `DATABASE_URL` in `docker-compose.yml` uses `db` as the hostname.
* **`sqlc` command not found**: Ensure `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest` was run successfully and your Go `bin` directory is in your system's PATH environment variable.
* **`curl` syntax errors**: For Windows CMD, use double quotes for the entire `-d` payload and escape internal double quotes (e.g., `\"`).

## License

This project is open-source and available under the [MIT License](LICENSE.md).
