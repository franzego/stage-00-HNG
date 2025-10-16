# Stage 00 Backend API 🚀

## Overview
This project implements a lightweight RESTful API built with Go and the Gin web framework. It features robust rate limiting using `tollbooth`, environment variable management via `godotenv`, and integrates with an external Cat Fact API to provide dynamic content.

## Features
-   **Gin Gonic**: Provides a high-performance HTTP web framework for building the API.
-   **Tollbooth**: Implements robust rate limiting to prevent API abuse and ensure stability.
-   **Godotenv**: Manages environment variables for flexible and secure configuration.
-   **net/http**: Facilitates external API integration for fetching dynamic content (e.g., Cat Facts).
-   **encoding/json**: Handles structured JSON responses for consistent and machine-readable output.
-   **os/signal, syscall**: Ensures graceful server shutdown upon receiving termination signals.

## Getting Started
To get this project up and running locally, follow these simple steps.

### Installation
🚀 Clone the repository to your local machine:
```bash
git clone https://github.com/franzego/stage-00-HNG.git
cd stage-00-HNG
```

📦 Ensure you have Go installed (version 1.25 or higher recommended). Then, fetch the project dependencies:
```bash
go mod tidy
```

### Environment Variables
⚙️ This project requires specific environment variables to function correctly. Create a `.env` file in the root directory of the project and populate it with the following:

-   `PORT`: The port number on which the server will listen.
    *Example:* `PORT=8080`
-   `catFactUrl`: The URL of the external Cat Fact API.
    *Example:* `catFactUrl=https://catfact.ninja/fact`

### Running the Application
▶️ Start the server by executing the `main.go` file:
```bash
go run main.go
```
The API will then be accessible at `http://localhost:[PORT]`.

## API Documentation
### Base URL
`http://localhost:[PORT]` (replace `[PORT]` with the value from your `.env` file, e.g., `8080`)

### Endpoints
#### GET /me
Retrieves a JSON object containing a predefined user profile, the current timestamp, and a random cat fact fetched from an external API.

**Request**:
No request body required.

**Response**:
```json
{
  "status": "success",
  "user": {
    "email": "davidenenama@gmail.com",
    "name": "Enenamah David",
    "stack": "Go/gin"
  },
  "timestamp": "2024-07-30T10:00:00Z",
  "fact": "A cat can jump up to six times its height."
}
```

**Errors**:
-   `502 Bad Gateway`: Occurs when the external Cat Fact API is unreachable, returns an invalid status code, or its response cannot be parsed. The response body will contain an error message detailing the issue.
    ```json
    {
      "error": "[CatApi] is unreachable"
    }
    ```
    _or_
    ```json
    {
      "error": "[CatAPI] returned status 500"
    }
    ```
    _or_
    ```json
    {
      "error": "[CatApi] JSON decode error: unexpected end of JSON input"
    }
    ```

## Technologies Used
This project leverages the following key technologies:

| Technology         | Description                                    | Link                                           |
| :----------------- | :--------------------------------------------- | :--------------------------------------------- |
| **Go**             | Primary programming language.                  | [golang.org](https://golang.org/)              |
| **Gin Gonic**      | High-performance HTTP web framework.           | [gin-gonic.com](https://gin-gonic.com/)        |
| **Tollbooth**      | Go middleware for rate-limiting HTTP requests. | [github.com/didip/tollbooth](https://github.com/didip/tollbooth) |
| **Godotenv**       | Loads environment variables from `.env` file.  | [github.com/joho/godotenv](https://github.com/joho/godotenv) |

## Author
**Enenamah David**
-   LinkedIn: [Your LinkedIn Profile](https://www.linkedin.com/in/davidenenamah)
-   Twitter: [@yourtwitterhandle](https://twitter.com/saint_franz)

## Badges
[![Go Version](https://img.shields.io/badge/Go-1.25-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Gin Framework](https://img.shields.io/badge/Gin_Gonic-v1.11.0-blue?style=flat&logo=go)](https://github.com/gin-gonic/gin)
[![Tollbooth Rate Limit](https://img.shields.io/badge/Tollbooth-v7.0.2-orange?style=flat)](https://github.com/didip/tollbooth)
[![Dotenv](https://img.shields.io/badge/GoDotEnv-v1.5.1-informational?style=flat)](https://github.com/joho/godotenv)
[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen?style=flat)](https://github.com/franzego/stage-00-HNG/actions)
