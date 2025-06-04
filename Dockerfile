# 1. Используем официальный образ Go
FROM golang:1.23-alpine

# 2. Устанавливаем git (нужен для go mod) и bash
RUN apk add --no-cache git

# 3. Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# 4. Копируем go.mod и go.sum, чтобы закешировать зависимости
COPY go.mod go.sum ./
RUN go mod download

# 5. Копируем остальной код
COPY . .

# 6. Собираем Go-приложение
RUN go build -o poison-bot ./cmd/bot

# 7. Запускаем приложение
CMD ["./poison-bot"]