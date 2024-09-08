# Inventory Management System

## Overview

This project is a comprehensive inventory management system built using a microservices architecture. The services included in this project are:

- **Inventory Service**: Manages inventory items and their quantities.
- **Notification Service**: Handles sending notifications related to inventory and order statuses.
- **Order Service**: Manages customer orders and order processing.
- **Shipment Service**: Tracks and updates shipment statuses as they progress through various stages.
- **User Service**: Manages users, their roles, and integrates role-based access control (RBAC). Handles user creation, deletion, and updates, with optional Kafka integration for broadcasting user-related events.
- **Auth Service**: Handles authentication and authorization, integrating with the user service for session-based authentication and token validation.

## Technologies Used

- **Frontend**: React with TypeScript
- **Backend**: Go
- **Database**: MongoDB
- **Message Broker**: Kafka

## Getting Started

### Prerequisites

- [Node.js](https://nodejs.org/) (for the frontend)
- [Go](https://golang.org/dl/) (for the backend services)
- [MongoDB](https://www.mongodb.com/try/download/community) (for the database)
- [Kafka](https://kafka.apache.org/downloads) (for messaging)

### Running the Application

1. **Frontend**:
   - Navigate to the `frontend` directory.
   - Install dependencies: `npm install`
   - Start the development server: `npm start`

2. **Backend Services**:
   - Navigate to each service directory (e.g., `inventory-service`, `notification-service`, `order-service`, `shipment-service`, `user-service`, `auth-service`).
   - Build and run the service:
     ```bash
     go mod tidy
     go run main.go
     ```

3. **MongoDB**:
   - Ensure MongoDB is running and accessible at `mongodb://localhost:27017`.

4. **Kafka**:
   - Ensure Kafka is running and accessible at `kafka:9092`.

### Docker

1. **Build Docker Images**:
   - Navigate to each service directory and build Docker images:
     ```bash
     docker build -t [service-name] .
     ```

2. **Run Docker Containers**:
   - Run the containers using Docker Compose or individual `docker run` commands.

## API Endpoints

### Inventory Service
- **POST /inventory**: Add a new inventory item.
- **GET /inventory**: Retrieve all inventory items.
- **GET /inventory/:id**: Retrieve a specific inventory item.

### Notification Service
- **POST /notifications**: Send a notification.

### Order Service
- **POST /orders**: Create a new order.
- **GET /orders**: Retrieve all orders.
- **GET /orders/:id**: Retrieve a specific order.

### Shipment Service
- **POST /shipments**: Create a new shipment.
- **GET /shipments**: Retrieve all shipments.
- **GET /shipments/:shipmentId**: Track the status of a specific shipment.

### User Service
- **POST /users**: Create a new user (Admins only).
- **GET /users**: Retrieve all users.
- **GET /users/:id**: Retrieve a specific user by ID.
- **PUT /users/:id**: Update user information (Admins only).
- **DELETE /users/:id**: Delete a user (Admins only).

### Auth Service
- **POST /auth/login**: Authenticate a user and obtain a session token.
- **POST /auth/logout**: End a user's session.
- **GET /auth/validate**: Validate a session token.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request to contribute.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](./LICENSE) file for details.



