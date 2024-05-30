# Start from the official Golang image to build the binary
FROM golang:1.22.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code from the current directory to the working directory inside the container
COPY . .

# Install necessary dependencies for CGO
RUN apt-get update && apt-get install -y gcc g++ make

# Build the Go app with CGO enabled
RUN CGO_ENABLED=1 GOOS=linux go build -o funda ./cmd/funda/

# Start a new stage from scratch
FROM alpine:latest

# Install necessary packages
RUN apk --no-cache add ca-certificates sqlite-libs

# Set the working directory inside the container
WORKDIR /root/

# Copy the pre-built binary file and configuration from the previous stage
COPY --from=builder /app/funda /app/app.yaml ./

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./funda"]
