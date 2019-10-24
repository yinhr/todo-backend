FROM golang:latest

ENV GO111MODULE=on
ENV ECHO_ENV=development
ENV DBHOST=host.docker.internal
ENV DBPORT=3306
ENV DBNAME=todo_db
ENV DBUSER=docker
ENV DBPASS=docker
ENV REDISHOST=redis
ENV REDISPORT=6379
ENV REDISPASS=secret

WORKDIR /go/src/todo-api

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

COPY . .
RUN go build -o /go/bin/app
RUN go get github.com/rubenv/sql-migrate/...
CMD ["sh", "/go/src/todo-api/run.sh"]

EXPOSE 4000
