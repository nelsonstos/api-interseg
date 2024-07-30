# Etapa de construcción: utiliza una imagen con Go instalado
FROM golang:1.22.5 AS builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia go.mod y go.sum para caché de dependencias
COPY go.mod go.sum ./

# Descarga las dependencias de Go
RUN go mod download

# Copia el código fuente del proyecto
COPY . .

# Compila el binario para Linux (sin CGO)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go

# Etapa final: usa una imagen más ligera para el contenedor final
FROM alpine:3.14

# Instala las dependencias necesarias, si las hay (libc6-compat, etc.)
RUN apk add --no-cache libc6-compat

# Establece el directorio de trabajo
WORKDIR /app

# Copia el binario compilado desde la etapa de construcción
COPY --from=builder /app/main .

# Copia el archivo .env al directorio de trabajo
COPY --from=builder /app/.env .

# Asegúrate de que el binario tenga permisos de ejecución
RUN chmod +x /app/main

# Expone el puerto en el que la aplicación está escuchando
EXPOSE 8080

# Comando por defecto para ejecutar la aplicación
CMD ["./main"]
