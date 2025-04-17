# Go-contacts app

Архетиктура приложения состоит из следующих микросервисов:
User Service, Subscriber service, Chat service, Auth service.

Для общения между сервисами будет использован протокол gRPC.
Для общения с клиентами будет использован REST API через Facade Gateway.
Для найстройки асинхронного общения между сервисами (между сервисом Auth и User)
будет использован брокер сообщений Kafka в связи с его надежностью и масштабируемостью.

## User Service

User Service - отвечает за регистрацию пользователя, редактирование профиля,
поиск пользователя по никнейму, получение данных пользователя.
Сервис подписывается на событие UserRegistered через брокер сообщений Kafka
от сервиса Auth и создавать нового пользователя.

API User Service содержит следующие методы:

Регистрация пользователя (Create) - POST /users

Редактирование профиля (Update) - PUT /users/{id}

Получение профиля по id (Get) - GET /users/{id}

Поиск пользователя по никнейму (Search) - GET /users/search

## Subscriber service

Subscriber service - отвечает за добавление пользователя в друзья,
удаление пользователя из друзей, подтверждение/отклонение запроса
на дружбу, просмотр списка своих друзей.

API Subscriber service содержит следующие методы:

Добавление пользователя в друзья (Make friend request) - POST /subscriber/request;

Удаление друга (Remove friend) - DELETE /subscriber

Подтверждение дружбы - (Confirm request) - POST /subscriber/accept;

Отклонение дружбы - (Decline request) - POST /subscriber/reject;

Получить список друзей пользователя - GET /subscriber/{id}

## Chat service

Chat service - отвечает за создание чата с другом, получение списка
чатов пользователя, отправку сообщения в чат и получения всех сообщений
пользователя из чата.

API Chat service содержит следующие методы:

Создание чата (Create Chat) - POST /chats

Получение списка чатов пользователя (List Chat) - GET /chats/{userID}

Отправку сообщения в чат (Send message to Chat) - POST /chats/{chatID}/messages

Получение всех сообщений пользователя из чата (Get Chat Messages) - GET /chats/{chatID}/messages

## Auth service

Auth service - отвечает за авторизацию пользователя.
Будет взаимодействовать с сервисом User по gRPC.
Будет публиковать события о регистрации нового пользователя UserRegistered, использую
при этом брокер сообщений Kafka. Это объясняется тем, что можно регистрировать и отвечать
клиенту быстро, а профиль создавать в фоне.

Регистрация пользователя - (Register) POST /auth/register

Логин, выдача токена - (Login) POST /auth/login

Проверка валидности токена - (Refresh token) GET /auth/verify

Выход, аннулирование токена - (Logout) POST /auth/logout

#

[![1.png](https://i.postimg.cc/qRTCF7X1/1.png)](https://postimg.cc/1gvtnsdF)
