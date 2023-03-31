# Utilisez une image Golang pour construire votre application
FROM golang:1.20 as builder

# Set the working directory
WORKDIR /app

# Copiez les fichiers nécessaires pour construire l'application
COPY . .

# Download the dependencies
RUN go mod download

# Compilez l'application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Utilisez une image légère pour exécuter l'application
FROM alpine:3.14

# Installez les certificats pour les requêtes HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiez l'exécutable
COPY --from=builder /app/main .

# Définissez la commande pour exécuter l'application
ENTRYPOINT ["./main"]
