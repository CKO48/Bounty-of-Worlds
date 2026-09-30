# ETAPA 1: Compilación
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copiar todo el código fuente del módulo
COPY . .

# Limpiar, descargar dependencias y generar/actualizar go.sum
RUN go mod tidy

# Compilar el ejecutable para Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# ETAPA 2: Imagen final superliviana
FROM alpine:latest

WORKDIR /app

# Instalar certificados CA para conexiones HTTPS (Discord API)
RUN apk --no-cache add ca-certificates

# Copiar el ejecutable generado en la etapa de compilación
COPY --from=builder /app/main .

CMD ["./main"]
