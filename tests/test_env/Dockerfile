FROM golang:latest

WORKDIR /matchme

COPY . .
COPY ../../wait-for-postgres.sh .

RUN ls -la

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash
RUN apt-get update
RUN apt-get install -y migrate

RUN chmod +x ./wait-for-postgres.sh

RUN go mod download

# Установка зависимости для тестов, если есть (например, тестовые библиотеки)
RUN go get -t ./...

# Сборка основного приложения
RUN go build -o main cmd/matchme/main.go

# Команда по умолчанию для выполнения тестов
CMD ["go", "test", "-v","./tests/..."]
