# User Management System

This project is a backend application for managing users.
**Backend**: Built with Go, handling the API, database connection, and user management logic.

## Prerequisites

Ensure you have the following installed:

- Go (version 1.18 or higher)
- Node.js (version 16 or higher)
- PostgreSQL
- Git
 
## Setup

### Backend Setup

1. Navigate to the backend folder and install dependencies:

    ```bash
    cd backend
    ```

2. Copy `.env` file from `.env.example` and update the variable values.

3. Install Go dependencies:

    ```bash
    go mod tidy
    ```

4. Run the backend application:

    ```bash
    go run main.go
    ```

    This will load environment variables, set up the database connection, and start the API routes.

