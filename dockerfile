FROM golang:alpine AS builder
RUN apk update && apk add --no-cache build-base
WORKDIR /app
COPY . .
RUN go mod download
RUN make build

FROM ubuntu:latest
RUN apt-get update && \
apt-get install -y curl
RUN apt-get install -y apache2 

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz | tar xvz
RUN mv migrate /usr/local/bin/migrate
WORKDIR /app

EXPOSE 8080

COPY --from=builder /app /app

RUN chmod +x entrypoint.sh

CMD ["./entrypoint.sh"]
