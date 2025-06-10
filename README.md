A simple CMS backend built with Go.

## Features

- RESTful API for content management
- Modular and extensible architecture
- SQLite/PostgreSQL support
- Basic authentication with register/login
- Forgot/reset password endpoints
- Google login stub
- Simple role and permission management

## Getting Started

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/golang-cms.git
    cd golang-cms
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Run the application:
    ```bash
    go run main.go
    ```

4. Run tests:
    ```bash
    go test ./...
    ```

## Configuration

Edit the `config.yaml` file to set up your database and environment variables.

## API Documentation

See [API.md](API.md) for detailed API endpoints and usage.

## License

This project is licensed under the MIT License.
