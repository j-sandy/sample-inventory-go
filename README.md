# E-commerce Inventory Backend API (Go)

A RESTful API for managing e-commerce inventory with session-based authentication, built with Go and an in-memory database. This project provides a simple yet functional backend for inventory management, demonstrating common API patterns and Go best practices.

## Features

*   **User Authentication**: Secure login and logout functionality using session tokens.
*   **Add Item**: Create new inventory items with details such as name, item-code, image, description, quantity, procurement date, manufacturing date, and expiry date.
*   **Update Item**: Modify details of existing inventory items.
*   **Delete Item**: Remove items from the inventory.
*   **Fetch Items**: Retrieve inventory items, with support for filtering by item-code, name, procurement date, and expiry date.

## Technologies Used

*   **Go**: The primary programming language for the backend.
*   **Gorilla Mux**: A powerful HTTP router and URL matcher for building Go web servers.
*   **In-memory Database**: A simple, thread-safe in-memory store for managing inventory items and user sessions (for demonstration purposes).
*   **UUID**: Used for generating unique session tokens.

## Project Structure

The project is organized into a `src` directory containing all Go source code, with a clear separation of concerns:

*   `./`: Project root directory.
    *   `README.md`: This file.
    *   `.gitignore`: Specifies intentionally untracked files to ignore.
    *   `src/`: Contains all Go source code.
        *   `src/main.go`: The application's entry point. It sets up the HTTP router, registers API handlers, and starts the server.
        *   `src/models/models.go`: Defines the Go structs for `Item`, `Session`, and `LoginRequest`, representing the data models used throughout the application.
        *   `src/store/store.go`: Implements the in-memory data storage logic for `Item` and `Session` objects, including CRUD operations and search functionality.
        *   `src/handlers/handlers.go`: Contains the HTTP handler functions responsible for processing incoming API requests, interacting with the `store`, and sending responses.
        *   `src/middleware/auth.go`: Provides the session authentication middleware, which validates session tokens for protected routes.
        *   `src/utils/response.go`: Offers utility functions for sending consistent JSON responses and error messages.

## Getting Started

### Prerequisites

Ensure you have Go installed on your system. Go version 1.22.12 or higher is recommended.

### Running the Server

1.  Navigate to the `src` directory:
    ```bash
    cd src
    ```
2.  Run the application:
    ```bash
    go run .
    ```
    The server will start on `http://localhost:8080`.

## API Endpoints

All API requests that modify or fetch inventory data require authentication via a `Bearer` token in the `Authorization` header.

### 1. Authenticate and Start Session

*   **Endpoint**: `POST /login`
*   **Description**: Authenticates a user and returns a session token.
*   **Request Body (JSON)**:
    ```json
    {
        "username": "admin",
        "password": "password"
    }
    ```
    *(Note: Credentials are hardcoded as "admin"/"password" for demonstration purposes.)*
*   **Success Response (200 OK)**:
    ```json
    {
        "message": "Login successful",
        "data": {
            "token": "e888d159-3c20-45d4-8947-ed0646cbd80c"
        }
    }
    ```

### 2. Logout and End Session

*   **Endpoint**: `POST /logout`
*   **Description**: Invalidates the current session token.
*   **Headers**:
    ```
    Authorization: Bearer <session_token>
    ```
*   **Success Response (200 OK)**:
    ```json
    {
        "message": "Logout successful"
    }
    ```

### 3. Add Item Details

*   **Endpoint**: `POST /items`
*   **Description**: Adds a new item to the inventory.
*   **Headers**:
    ```
    Authorization: Bearer <session_token>
    Content-Type: application/json
    ```
*   **Request Body (JSON)**:
    ```json
    {
        "item_code": "ITEM001",
        "name": "Laptop",
        "image": "laptop.jpg",
        "description": "High performance laptop",
        "quantity": 10,
        "procurement_date": "2023-01-15",
        "manufacturing_date": "2023-01-01",
        "expiry_date": "2028-01-01"
    }
    ```
*   **Success Response (201 Created)**:
    ```json
    {
        "message": "Item added successfully",
        "data": {
            "item_code": "ITEM001",
            "name": "Laptop",
            "image": "laptop.jpg",
            "description": "High performance laptop",
            "quantity": 10,
            "procurement_date": "2023-01-15",
            "manufacturing_date": "2023-01-01",
            "expiry_date": "2028-01-01"
        }
    }
    ```

### 4. Update Item Details

*   **Endpoint**: `PUT /items/{itemCode}`
*   **Description**: Updates an existing item's details based on its `itemCode`.
*   **Headers**:
    ```
    Authorization: Bearer <session_token>
    Content-Type: application/json
    ```
*   **Request Body (JSON)**:
    ```json
    {
        "item_code": "ITEM001",
        "name": "Gaming Laptop",
        "image": "gaming_laptop.jpg",
        "description": "High performance gaming laptop with RGB",
        "quantity": 8,
        "procurement_date": "2023-01-15",
        "manufacturing_date": "2023-01-01",
        "expiry_date": "2028-01-01"
    }
    ```
*   **Success Response (200 OK)**:
    ```json
    {
        "message": "Item updated successfully",
        "data": {
            "item_code": "ITEM001",
            "name": "Gaming Laptop",
            "image": "gaming_laptop.jpg",
            "description": "High performance gaming laptop with RGB",
            "quantity": 8,
            "procurement_date": "2023-01-15",
            "manufacturing_date": "2023-01-01",
            "expiry_date": "2028-01-01"
        }
    }
    ```

### 5. Delete Item

*   **Endpoint**: `DELETE /items/{itemCode}`
*   **Description**: Deletes an item from the inventory based on its `itemCode`.
*   **Headers**:
    ```
    Authorization: Bearer <session_token>
    ```
*   **Success Response (200 OK)**:
    ```json
    {
        "message": "Item deleted successfully"
    }
    ```

### 6. Fetch Item Details

*   **Endpoint**: `GET /items`
*   **Description**: Retrieves a list of inventory items. Supports filtering by query parameters.
*   **Headers**:
    ```
    Authorization: Bearer <session_token>
    ```
*   **Query Parameters (Optional)**:
    *   `item-code`: Filter by a specific item code.
    *   `name`: Filter by item name.
    *   `procurement-date`: Filter by procurement date.
    *   `expiry-date`: Filter by expiry date.
*   **Example Request**:
    ```bash
    curl -X GET "http://localhost:8080/items?name=Laptop&procurement-date=2023-01-15" \
         -H "Authorization: Bearer <session_token>"
    ```
*   **Success Response (200 OK)**:
    ```json
    {
        "message": "Items fetched successfully",
        "data": [
            {
                "item_code": "ITEM001",
                "name": "Gaming Laptop",
                "image": "gaming_laptop.jpg",
                "description": "High performance gaming laptop with RGB",
                "quantity": 8,
                "procurement_date": "2023-01-15",
                "manufacturing_date": "2023-01-01",
                "expiry_date": "2028-01-01"
            }
        ]
    }
    ```
    *(Returns an empty array if no items match the criteria or if the inventory is empty.)*

## Authentication

All endpoints except `/login` require a valid session token. After a successful login, the API returns a `Bearer` token. This token must be included in the `Authorization` header of subsequent requests in the format: `Authorization: Bearer <your_session_token>`.

## Error Handling

The API provides consistent JSON error responses for various scenarios (e.g., invalid payload, unauthorized access, item not found).

*   **Error Response Format**:
    ```json
    {
        "error": "Error message describing the issue"
    }
