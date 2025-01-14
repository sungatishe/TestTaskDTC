# Устанавливаем базовый образ
FROM golang:1.22.4-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем зависимости и устанавливаем их
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Сборка исполняемого файла
RUN go build -o /order ./cmd/order/main.go



# Создаем минимальный образ для запуска
FROM alpine:latest
WORKDIR /root/

# Копируем скомпилированное приложение из builder'а
COPY --from=builder /order .

# Копируем конфигурационный файл
COPY config/config.yaml ./config.yaml

COPY docs ./docs

# Экспорт порта
EXPOSE 8080

# Команда для запуска приложения
CMD ["./order"]
