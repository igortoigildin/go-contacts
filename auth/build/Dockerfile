# MULTI STAGE EXAMPLE

#######################################
# STAGE 1. BUILD STAGE
#######################################

# Для сборки нашего приложения на go требуется образ ОС, в котором установлен golang нужной нам версии.
# Alpine выбран из-за его небольшого размера по сравнению с Ubuntu.
FROM golang:1.24.2-alpine AS builder

# Устанавливаем место назначения для COPY
WORKDIR /app

# Копируем файлы go.mod и go.sum в WORKDIR
COPY go.mod go.sum ./
# Скачиваем необходимые Go модули (зависимости нашего проетка)
RUN go mod download

# Копируем все исходные go файлы нашего проекта в образ
# https://docs.docker.com/reference/dockerfile/#copy
COPY . .
# Собираем бинарный файл нашего приложения
RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service ./cmd

#######################################
# STAGE 2. FINAL STAGE
#######################################

FROM alpine:3.19

WORKDIR /

# Копируем бинарник из этапа сборки
COPY --from=builder /app/auth-service /auth-service

# Устанавливаем рабочую директорию и порт
WORKDIR /

EXPOSE 8080

# Указываем команду запуска
CMD ["/auth-service"]