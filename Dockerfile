# Этап сборки
FROM golang:1.24-alpine AS builder

# Устанавливаем рабочую директорию в контейнере
WORKDIR /app

# Копируем все файлы Go в рабочую директорию контейнера
COPY . .

# Скачиваем зависимости и выполняем сборку
RUN go mod tidy
RUN go build -o main ./cmd

# Этап для работы приложения
FROM alpine:latest

# Устанавливаем необходимые пакеты (например, для работы с сертификатами)
RUN apk --no-cache add ca-certificates

# Копируем собранное приложение из предыдущего этапа
WORKDIR /root/
COPY --from=builder /app/main .

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]
