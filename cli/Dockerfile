# Utilisez une image Golang pour construire votre application
FROM golang:1.20 as builder

# Set the working directory
WORKDIR /app

# Copiez les fichiers nécessaires pour construire l'application
COPY . ./

# Download the dependencies
RUN go mod download

# Build the app
RUN CGO_ENABLED=0 GOOS=linux go build -o cli main.go commands.go

# Second stage: create a minimal image to deploy the app
FROM alpine:3.14

# Copy the binary from the builder stage
COPY --from=builder /app/cli /usr/local/bin/cli

# Set the app as the entrypoint
ENTRYPOINT ["cli"]