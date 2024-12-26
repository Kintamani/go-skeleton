FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk update
RUN apk add --no-cache \
    bash \
    curl \
    git 

RUN go install github.com/air-verse/air@v1.60.0
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/go-task/task/v3/cmd/task@v3.40.1

COPY . .
RUN go mod download

CMD [ "air", "-c", ".air.toml" ]