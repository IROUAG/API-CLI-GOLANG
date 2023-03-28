# Use the official Go image as the base
FROM golang:1.17

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port the app will run on
EXPOSE 8080

# Start the application
CMD ["./main"]