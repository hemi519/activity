# Start from the Go base image
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module file
COPY go.mod ./

# Download Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o /app/cmd/main ./cmd

# Set the entrypoint for the container
ENTRYPOINT ["./cmd/main"]
