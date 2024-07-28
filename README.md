# SuperTodo

SuperTodo is a simple task management application built with Wails, using SQLite for data storage and Go for the backend.

## Features

- Add, retrieve, update, and remove tasks.
- Cross-platform support 

### Dependencies

- [Wails CLI](https://wails.io/)
- [go-sqlite3](https://github.com/mattn/go-sqlite3)
- [sqlc](https://github.com/kyleconroy/sqlc)
- [dbmate](https://github.com/amacneil/dbmate)

## Getting Started

### Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/lakubudavid/todo-wails.git
   cd todo-wails
   ```

2. **Install dependencies:**

   Make sure you have the Wails CLI installed:

   ```sh
   go install github.com/wailsapp/wails/v2/cmd/wails@latest
   ```
   As well as `dbmate` (for database migrations) and sqlc
   ```sh
   go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest  
   ```

### Database Setup

1. **Create Database:**

   Ensure the `todo.db` file is located in `~/todo-wails/`.

2. **Apply Database Schema:**

   The schema will be applied automatically on the first run of the application. If needed, the schema is located in `db/schema.sql`.

### Running the Application

1. **Development Mode:**

   ```sh
   just dev
   ```

2. **Build the Application:**

   ```sh
   just build
   ```

### Project Structure

- `db`: Database schema and SQL queries.
- `frontend`: Frontend assets.


### Contributing

Contributions are welcome! Please open an issue or submit a pull request.

### License

This project is licensed under the MIT License. See the `LICENSE` file for details.
