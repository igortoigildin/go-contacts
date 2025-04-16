# Go-contacts app

Архетиктура приложения состоит из следующих микросервисов:
User Service, Subscriber service, Chat service, Auth service.

#### User Service

User Service - отвечает за регистрацию пользователя, редактирование профиля,
поиск пользователя по никнейму, получение данных пользователя.

API User Service содержит следующие методы:

Регистрация пользователя (Create) - POST /users

Редактирование профиля (Update) - PUT /users/{id}

Получение профиля по id (Get) - GET /users/{id}

Поиск пользователя по никнейму (Search) - GET /users/search

#### Subscriber service

Subscriber service - отвечает за добавление пользователя в друзья,
удаление пользователя из друзей, подтверждение/отклонение запроса
на дружбу, просмотр списка своих друзей.

API Subscriber service содержит следующие методы:

Добавление пользователя в друзья (Make friend request) - POST /subscriber/request;

Удаление друга (Remove friend) - DELETE /subscriber

Подтверждение дружбы - (Confirm request) - POST /subscriber/accept;

Отклонение дружбы - (Decline request) - POST /subscriber/reject;

Получить список друзей пользователя - GET /subscriber/{id}

#### Chat service

Chat service - отвечает за создание чата с другом, получение списка
чатов пользователя, отправку сообщения в чат и получения всех сообщений
пользователя из чата.

API Chat service содержит следующие методы:

Создание чата (Create Chat) - POST /chats

Получение списка чатов пользователя (List Chat) - GET /chats/{userID}

Отправку сообщения в чат (Send message to Chat) - POST /chats/{chatID}/messages

Получение всех сообщений пользователя из чата (Get Chat Messages) - GET /chats/{chatID}/messages

#### Auth service

Auth service - отвечает за авторизацию пользователя.

Логин, выдача токена - (Login) - POST /auth/login

Проверка валидности токена - (Refresh token) - GET /auth/verify

Выход, аннулирование токена - (Logout) - POST /auth/logout

[![Screenshot-2025-04-15-at-1-09-54-PM.png](https://i.postimg.cc/5tL8tNdv/Screenshot-2025-04-15-at-1-09-54-PM.png)](https://postimg.cc/NyGKJt0f)
