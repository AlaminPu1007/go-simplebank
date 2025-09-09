# Build Stage 
FROM golang:1.25.1-alpine3.22 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.22
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh .
COPY wai-for.sh .
COPY db/migrations ./db/migrations

EXPOSE 8000 
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]