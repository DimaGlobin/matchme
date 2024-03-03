FROM golang:latest

WORKDIR /matchme

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash
RUN apt-get update
RUN apt-get install -y migrate

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o main cmd/matchme/main.go

CMD ["./main"]