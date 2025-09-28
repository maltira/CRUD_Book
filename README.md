
# CRUD_Book

Учебный pet-проект — REST API для управления книгами, реализованный на Go + Gin + GORM + PostgreSQL. Поддерживает создание, получение, обновление и удаление книг (CRUD).
Документация API генерируется автоматически через Swagger.


## Стек технологий

- [Go](https://go.dev/)
- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [godotenv](https://github.com/joho/godotenv)
- [swag](https://github.com/swaggo/swag)
## Установка

Клонируем репозиторий
```bash
  git clone https://github.com/maltira/CRUD_Book.git
  cd crud_book
```

Настраиваем окружение (.env)
```bash
  APP_PORT=<PORT>
  DB_HOST=<HOST>
  DB_PORT=<PORT>
  DB_USER=<USER>
  DB_PASSWORD=<PASS>
  DB_NAME=<NAME>
```

Устанавливаем зависимости
```bash
  go mod tidy
```

Генерируем Swagger
```bash
  swag init -g cmd/main.go
```

Запускаем сервер
```bash
  go run cmd/main.go
```
