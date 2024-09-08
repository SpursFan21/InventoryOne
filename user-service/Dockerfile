# Use official Golang image as the base image
FROM golang:1.20-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o user-service .

# Expose port 8086 to the outside world
EXPOSE 8086

# Command to run the executable
CMD ["./user-service"]
