# Use the official Go image as the base
FROM golang:golang:1.20

# Set the working directory
WORKDIR /app

# Copy the entire project to the working directory
COPY . ./

# Download the dependencies
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose the port the app will run on
EXPOSE 8080

# Start the application
CMD ["./main"]