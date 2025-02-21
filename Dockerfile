# Etapa de construcción
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o server ./server/main.go

# Etapa final (más liviana)
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
RUN chmod +x /root/server  # Permisos de ejecución
RUN apk add --no-cache libc6-compat  # Evita problemas con Alpine
EXPOSE 8080
CMD ["./server"]