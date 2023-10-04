# BWG
Тестовое задание

## :open_file_folder: Структура репозитория

- `cmd/main.go` - главный файл, в котором запускается сервер http
- `config` - папка с конфигом
- `entity` - папка с сущностями
- `internal` - папка с кодом проекта
    - `handler` - папка с rest ручками
        - `handler.go` - файл с ручками
        - `response.go` - файл со структурой respons
        - `account.go` - файл с ручками для сущности account
        - `currency.go` - файл с кодом, который отправляет запрос на api для проверка валидности кода валюты
        - `transaction.go` - файл с ручками для сущности transaction
        - `user.go` - файл с ручками для сущности user
    - `repository` - папка с уровнем storage
        - `cache.go` - файл с методами для кэширования Redis
        - `postgres.go` - бд Postgres
        - `redis.go` - бд Redis
        - `repository.go` - файл с интерфесами для репозитория
        - `user.go` - методы для бд сущности user
        - `account.go` - методы для бд сущности account
        - `transaction.go` - методы для бд сущности transaction
    - `service` - папка с бизнес логикой
        - `service.go` - файл с интерфейсами для сервиса
        - `user.go` - бизнес логика для user
        - `account.go` - бизнес логика для account
        - `transaction.go` - бизнес логика для transaction
- `migrations` - папка с миграциями, `Dockerfile`-ом и `.sh`-скриптом для их накатывания на БД
- `Makefile` - обычный Makefile с различными командами для запуска сервера
- `docker-compose.yaml` - docker-compose для запуска контейнеров(postgres, redis, api)
- `Dockerfile` - Dockerfile для запуска сервера в докер-контейнере

## О проекте

Были использованы 3 сущности(User, Account, Transaction)

User
```
    id         SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL ,
    surname  VARCHAR(255) NOT NULL
```
Account - счет пользователя
```
    id             SERIAL PRIMARY KEY,
    currency_code  VARCHAR(3),
    active_balance FLOAT DEFAULT 0,
    frozen_balance FLOAT DEFAULT 0,
    user_id        INT,
    FOREIGN KEY (user_id) REFERENCES users (id)
```
Transaction - транзакции
```
    id                   SERIAL PRIMARY KEY,
    currency_code        VARCHAR(3),
    transaction_status   VARCHAR(25),
    sender_account_id    INT,
    recipient_account_id INT,
    amount               FLOAT,
    transaction_date     TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (sender_account_id) REFERENCES accounts (id),
    FOREIGN KEY (recipient_account_id) REFERENCES accounts (id)
```
 # Ручки

POST /createUser - ручка создания юзера

JSON

```
  "username": "name",
  "surname": "sur"
```

GET /getUser/:id - получение юзера по id

GET /getUsers - получение списка юзеров с пагинацией(По дэфолту /getUsers?per_page=10&page=1)

POST /createAccount - создание счета юзера

JSON

```
  "currency_code": "RUB",
  "user_id": 1
```

GET /getUserAccounts/:id - полчучение всех счетов юзера по id юзера

POST /invoice - начисление на счет

JSON

```
  "currency_code": "RUB",
  "recipient_account_id": 1, - номер счета получателя
  "amount": 10 - сумма
```

POST /withdraw

JSON

```
  "currency_code": "RUB",
  "sender_account_id": 2, - номер счета отправителя
  "recipient_account_id": 1, - номер счета получателя(опционально)
  "amount": 10 - сумма
```

POST /initTransaction/:id - ручка инициализации транзакции по id транзакции(по дэфолту все транзакции имеют статус Created или Error, если произошла ошибка)

GET /getTransaction/:id - получение транзакции по id

GET /getTransactions - полчуение списка транзакций с пагинацией(По дэфолту /getTransactions?per_page=10&page=1)

GET /getAccountTransactions/:id - ручка получения транзакций по счету

## :hammer: Запуск

docker compose up --build

make migrations-up
(__необходимо установить go migrate__)
