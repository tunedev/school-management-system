# Use a multi-stage build to compile the Go application
FROM golang:1.17.5 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the Go application source code
COPY . .

# Build the Go application for Linux/amd64
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o main cmd/server/main.go

# Use a minimal Docker image to run the Go application
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build /app/main /main
CMD ["/main"]
