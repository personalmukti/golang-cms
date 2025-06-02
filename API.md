# API Documentation

This document describes the RESTful API endpoints for the golang-cms backend.

## Authentication

All protected endpoints require a valid JWT token in the `Authorization` header.

### Login

- **POST** `/api/auth/login`
  - Request Body:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - Response:
    ```json
    {
      "token": "jwt_token"
    }
    ```

---

## Users

### Register a New User

- **POST** `/api/users/register`
  - Request Body:
    ```json
    {
      "username": "string",
      "password": "string",
      "email": "string"
    }
    ```
  - Response: User object

### Get Current User

- **GET** `/api/users/me`
  - Headers: `Authorization: Bearer <token>`
  - Response: User object

---

## Content

### List All Content

- **GET** `/api/content`
  - Response: Array of content objects

### Get Content by ID

- **GET** `/api/content/{id}`
  - Response: Content object

### Create Content

- **POST** `/api/content`
  - Headers: `Authorization: Bearer <token>`
  - Request Body:
    ```json
    {
      "title": "string",
      "body": "string",
      "status": "published|draft"
    }
    ```
  - Response: Content object

### Update Content

- **PUT** `/api/content/{id}`
  - Headers: `Authorization: Bearer <token>`
  - Request Body: Same as create
  - Response: Updated content object

### Delete Content

- **DELETE** `/api/content/{id}`
  - Headers: `Authorization: Bearer <token>`
  - Response: Success message

---

## Error Handling

All errors return a JSON object:
```json
{
  "error": "Error message"
}