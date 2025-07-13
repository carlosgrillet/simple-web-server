# Simple Web Server

This is a minimalistic web server written in Go (Golang). It listens on port `8080` and responds with "ok" to all incoming HTTP requests at the root endpoint (`/`).

## Features
- Lightweight and simple implementation.
- Responds with "ok" to all requests at the root endpoint.
- Dockerized for easy deployment.

## Prerequisites
To build and run this project locally, you need:
- [Go](https://golang.org/) (version 1.23.3 or higher)
- [Docker](https://www.docker.com/) (optional, for containerized deployment)

## Project Structure
The project consists of the following files:
- `main.go`: The main Go source file containing the web server logic.
- `go.mod`: The Go module file defining dependencies and the Go version.
- `Dockerfile`: Defines the multi-stage Docker build process for the application.
- `docker-compose.yml`: Simplifies running the application using Docker Compose.

## Running Locally
### Without Docker
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd simple-web-server
   ```
2. Build and run the application:
   ```bash
   go run main.go
   ```
3. Access the server:
   Open your browser or use `curl` to access `http://localhost:8080`. You should see the response:
   ```
   ok
   ```

### With Docker
1. Build the Docker image:
   ```bash
   docker build -t go-web-server .
   ```
2. Run the container:
   ```bash
   docker run -p 8080:8080 go-web-server
   ```
3. Access the server:
   Open your browser or use `curl` to access `http://localhost:8080`.

### Using Docker Compose
1. Start the server using Docker Compose:
   ```bash
   docker-compose up --build
   ```
2. Access the server:
   Open your browser or use `curl` to access `http://localhost:8080`.

## Customization
- **Port**: To change the port, modify the `http.ListenAndServe` function in `main.go` and update the `ports` section in `docker-compose.yml`.
- **Response**: Customize the response by editing the `root` function in `main.go`.

## License
This project is open-source and available under the [MIT License](LICENSE).
