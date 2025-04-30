# Собрать все и поднять
all: build up

# Билдим все сервисы через docker-compose
build:
	docker compose build

# Поднимаем в фоне
up:
	docker compose up -d

# Останавливаем и удаляем контейнеры
down:
	docker compose down

# Перезапуск
restart: down up
