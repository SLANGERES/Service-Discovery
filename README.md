# 🌐 Lightweight Service Discovery in Go

This is a simple and efficient service discovery system built from the ground up in Go. It provides a straightforward REST API for services to register themselves and for clients to find them. The system is designed to be lightweight, fast, and easy to integrate into microservice architectures.

## ✨ Features

- **Service Registration**: A POST endpoint for services to register their name, host, and port with the discovery service.
- **Service Discovery**: A GET endpoint for clients to find the available instances of a registered service.
- **Health Checks**: Uses a Time-to-Live (TTL) mechanism to automatically deregister services that have not sent a heartbeat.
- **In-Memory Storage**: Provides a fast and simple storage solution for service data.

## 🚀 Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You need to have Go installed on your machine.

- Go (version 1.20 or higher)

### Installation

Clone the repository to your local machine.

```bash
git clone https://github.com/your-username/Service-Discovery.git
cd Service-Discovery
```

### Running the Service

The main application is located in the cmd directory.

```bash
go run cmd/main.go
```

The server will start and listen on http://localhost:9000 by default.

## 💻 API Usage

Interact with the service discovery registry using the following REST API endpoints.

### 1. Register a Service

- **Endpoint**: `POST /register`
- **Description**: Registers a new service or updates an existing one. The TTL (Time-to-Live) is automatically set by the server.
- **Request Body**:
  ```json
  {
    "name": "my-service",
    "host": "192.168.1.10",
    "port": 8080
  }
  ```
- **Example cURL**:
  ```bash
  curl -X POST http://localhost:9000/register \
  -H "Content-Type: application/json" \
  -d '{"name": "my-service", "host": "192.168.1.10", "port": 8080}'
  ```

### 2. Discover Services

- **Endpoint**: `GET /discover/:serviceName`
- **Description**: Retrieves a list of all currently active instances for a given service.
- **Path Parameter**: `serviceName` (e.g., my-service)
- **Example cURL**:
  ```bash
  curl -X GET http://localhost:9000/discover/my-service
  ```
- **Example Response**:
  ```json
  [
    {
      "name": "my-service",
      "host": "192.168.1.10",
      "port": 8080,
      "ttl": 60
    },
    {
      "name": "my-service",
      "host": "192.168.1.11",
      "port": 8080,
      "ttl": 60
    }
  ]
  ```

## 📂 Project Structure

- `cmd/main.go`: The entry point for the application.
- `internal/api/`: Contains the Gin handlers for the REST API.
- `internal/models/`: Defines the data structures (e.g., Service).
- `internal/storage/`: Manages the in-memory storage of services.

## 👋 Contributing

Contributions are welcome! Please feel free to open a pull request or submit an issue on the repository.

## 📜 License

This project is licensed under the MIT License - see the LICENSE file for details.