# User Management System

This project is a full-stack application for managing users. It consists of two main parts:
1. **Backend**: Built with Go, handling the API, database connection, and user management logic.
2. **Frontend**: A React application with Vite, Axios, Formik, and Tailwind CSS, providing a user interface for interacting with the backend.

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

### Frontend Setup

1. Navigate to the frontend folder:

    ```bash
    cd frontend
    ```

2. Copy `.env` file from `.env.example` and update the variables.

3. Install the Node.js dependencies:

    ```bash
    npm install
    ```

4. Start the frontend development server:

    ```bash
    npm run dev
    ```

    This will start the frontend application on a local development server.

5. To build the production version of the frontend, run:

    ```bash
    npm run build
    ```

5. To preview the production build:

    ```bash
    npm run preview
    ```
