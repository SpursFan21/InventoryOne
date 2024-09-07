# Inventory Management System

## Overview

This project is a comprehensive inventory management system built using a microservices architecture. The services included in this project are:

- **Inventory Service**: Manages inventory items and their quantities.
- **Notification Service**: Handles sending notifications related to inventory and order statuses.
- **Order Service**: Manages customer orders and order processing.
- **Shipment Service**: Tracks and updates shipment statuses as they progress through various stages.

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
   - Navigate to each service directory (e.g., `inventory-service`, `notification-service`, etc.).
   - Build and run the service:
     ```bash
     go mod tidy
     go run main.go
     ```

3. **MongoDB**:
   - Ensure MongoDB is running and accessible at `mongodb://localhost:27017`.

4. **Kafka**:
   - Ensure Kafka is running and accessible at `kafka:9092`.

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

## Contributing

Contributions are welcome! Please open an issue or submit a pull request to contribute.

## License

This project is licensed under the MIT License.

