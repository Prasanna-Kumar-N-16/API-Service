# Use the official Golang image as the base image
FROM golang:1.18-alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy the source code to the container
COPY . .

# Build the Go app
RUN go build -o api-service .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
