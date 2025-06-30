# Study Case: Basic Load Balancing

This study case dives into the basics of load balancing using Nginx and a Go-based backend. It demonstrates how requests can be distributed across multiple backend instances to improve scalability and reliability.

## Objectives

- Learn the fundamentals of load balancing in distributed systems.
- Set up Nginx as a load balancer.
- Use Docker and Docker Compose for containerized deployments.
- Observe how a simple Go-based backend API behaves under load balancing.

## Project Structure

- `main.go`: Go code for the backend API.
- `Dockerfile`: Docker image definition for the backend API.
- `docker-compose.yml`: Configuration for the backend services and Nginx load balancer.
- `nginx.conf`: Nginx configuration file.
- `go.mod`: Go module file.

## Getting Started

1. Make sure Docker and Docker Compose are installed.
2. Start the services with:

   ```bash
   docker-compose up
   ```

3. Open your browser and go to `http://localhost:8080` to interact with the load balancer.

## Backend API

The backend API is a simple Go application that responds with "Hello world!" to GET requests at the root endpoint (`/`).

## Load Balancer

Nginx acts as the load balancer, routing incoming requests to one of the backend API instances (`api1`, `api2`, `api3`). 

It also adds a custom header `X-Upstream-Addr` to show which backend instance handled the request.