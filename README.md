# Task Management API

This is a simple **Task Management API** built using **Go (Golang)** and **Gin** framework. It allows users to create, view, update, and delete tasks. The API is backed by a **PostgreSQL** database, and it supports both **public** and **protected** endpoints.

## Features

- **Task management:** CRUD operations for tasks.
- **Authentication:** Token-based authentication for protected routes.
- **Database Integration:** PostgreSQL as the database with **GORM** ORM for database management.
- **Task Status:** Tasks can be in three states: `pending`, `in-progress`, and `completed`.
- **Migration Support:** Uses **golang-migrate** to manage database migrations.

## Technologies

- **Go (Golang)**: The backend language.
- **Gin**: Web framework for building RESTful APIs in Go.
- **PostgreSQL**: Relational database for storing tasks.
- **GORM**: ORM for interacting with the PostgreSQL database.
- **JWT**: Token-based authentication for protected routes.
- **golang-migrate**: Database migration tool.

## Endpoints

### Public Endpoints:
These routes do not require authentication.

- `GET /public/tasks`: Retrieve a list of all tasks.

### Protected Endpoints:
These routes require a valid **JWT token** for access.

- `GET /tasks/{id}`: Retrieve a task by its ID.
- `POST /tasks`: Create a new task.
- `PUT /tasks/{id}`: Update an existing task.
- `DELETE /tasks/{id}`: Delete a task.

## Database Setup

1. Create a `.env` file at the root of the project and add your PostgreSQL database URL:

   `DATABASE_URL=postgres://username:password@localhost:5432/dbname?sslmode=disable`

## Running the Application

- Clone the Repo
    `git clone git@github.com:guptachhavi27/task-api.git`
- Navigate into the project directory:
    `cd task-api`
- Make sure you have Go and PostgreSQL installed.
- Create a .env file with the following content:
    `DATABASE_URL=postgres://username:password@localhost:5432/dbname?sslmode=disable`
    Replacing username, password and dbname as needed.
- Run the application:
    `go run main.go`

## Testing the application
    - The postman collection to test this application can be found at
    `https://www.postman.com/guptachhavi27/task-api/collection/wo8ups7/task-management-api?share=true`

## Pagination
- The public GET /public/tasks endpoint supports pagination using the following query parameters:
    page: Page number (default: 1).
    page_size: Number of tasks per page (default: 10).
- Example Request:
 `GET /public/tasks?page=2&page_size=5`