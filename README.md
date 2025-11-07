# Distributed File Storage

A simple distributed file storage system. This project demonstrates the basics of storing, retrieving, and managing files across multiple nodes or services.

## Features

- Store files in a distributed manner.
- Retrieve files by unique IDs.
- Basic replication and redundancy.
- Simple REST API for interacting with the system.

## Getting Started

### Prerequisites

- Go 1.19+ (for backend/server)
- Docker & Docker Compose (optional, for running services in containers)

### Running the Application

1. **Clone the repository:**

   ```bash
   git clone <repository-url>
   cd distributed-file-storage
   ```

2. **Start the services:**

   ```bash
   go run main.go
   ```

   Or to use Docker Compose:

   ```bash
   docker-compose up --build
   ```

3. **Use the API:**

   - Upload a file: `POST /upload`
   - Download a file: `GET /files/{id}`
   - List files: `GET /files`

## Project Structure

- `main.go` - Entry point for the application
- `/internal` - Core application logic
- `/api` - REST API handlers
- `/storage` - Storage engine implementation

## Configuration

You may configure ports and storage directories via environment variables or a `.env` file.

## License

MIT License

