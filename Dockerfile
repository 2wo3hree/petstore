# Используем минимальный образ Golang
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Копируем только файлы для зависимости и устанавливаем их (кешируется!)
COPY go.mod go.sum ./
RUN go mod tidy && go mod download

# Копируем исходный код (изменение здесь не ломает кэш `go mod download`)
COPY . .

RUN go mod tidy

# Компилируем бинарник
RUN go build -o server ./cmd/web/main.go

# Используем минимальный базовый образ
FROM alpine:latest

WORKDIR /root/

# Устанавливаем зависимости
RUN apk --no-cache add ca-certificates postgresql-client

# Копируем скомпилированный бинарник (без лишних файлов)
COPY --from=builder /app/server .

# Копируем файл .env
COPY .env ./

COPY --from=builder /app/internal/db/migrations /root/internal/db/migrations

# wait-for-postgres
COPY wait-for-postgres.sh .
RUN chmod +x wait-for-postgres.sh

ENTRYPOINT ["./wait-for-postgres.sh"]

# Открываем порт
EXPOSE 8080

# Запускаем сервер
CMD ["./server"]