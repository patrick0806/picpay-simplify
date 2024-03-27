# PicPay Simplify

## Project Status: In Development
This project is based in the [picpay tech challange](https://github.com/PicPay/picpay-desafio-backend)
I want to made this in go for improve my skills in language and implement other features like concurence and safe transafer with concurence

## Api Development Process

To running this project in dev mode you need follow this steps

### 1 - Create and running local database
```bash
    docker run --name picpay-postgres \
    -p 5432:5432 \
    -e POSTGRES_DB=picpay \
    -e POSTGRES_USER=root \
    -e POSTGRES_PASSWORD=123 \
    -d postgres:14.4-alpine
```

### 2 - Running migrations <br>
 - Install migrate cli [see how.](https://github.com/golang-migrate/migrate)
 - Install make command [see how.](https://pt.linux-console.net/?p=14595#:~:text=Como%20instalar%20e%20usar%20o%20Make%20no%20Ubuntu,de%20comando%20make%2C%20execute%20o%20comando%3A%20%24%20make--version)
 - running migrations : 
 ```bash
    make migrate
 ```
 - create migration : 
 ```bash
    make createmigration
 ```
  - drop migrations : 
 ```bash
    make migratedown
 ```

### 3 - Install Dependencies
```bash
    go mod tidy
```

### 4 - Running App
```bash
    go run cmd/rest/main.go
```

## Reflection

I'm create this project for learn more about go, and your bases
for this project i don't pretend use any framework for understand more about golang
Maybe i use some libs for make things like Swagger, Router like mux or things like that