FROM ubuntu:latest
RUN apt-get update && \
    apt-get install -y curl golang gcc make mysql-client

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz | tar xvz
RUN mv migrate /usr/local/bin/migrate
WORKDIR /app
COPY . .
EXPOSE 8080
RUN go mod download
RUN make build

# RUN migrate -path database/migration/ -database "mysql://root:password@tcp(db:3306)/LMS" up

# CMD ["./build/MVC-LMS-SDS"]
RUN chmod +x entrypoint.sh
CMD ["./entrypoint.sh"]
