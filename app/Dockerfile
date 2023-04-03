# Utilisez une image Golang pour construire votre application
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copiez les fichiers nécessaires pour construire l'application
COPY . ./

# Download the dependencies
RUN go mod download

# Build the Go app
RUN go build -o app .

# Expose the port the app will run on
EXPOSE 8080

# Start the application
CMD ["./app"]