FROM golang:1.24-alpine AS build

WORKDIR /app

# Copiar archivos de dependencias desde apps/api
COPY apps/api/go.mod apps/api/go.sum ./
RUN go mod download

# Copiar el código fuente y compilar
COPY apps/api/ ./
RUN CGO_ENABLED=0 go build -o /app/main .

# Runtime ligero
FROM alpine:3.21
WORKDIR /app
COPY --from=build /app/main /app/main

EXPOSE 8080
CMD ["/app/main"]
