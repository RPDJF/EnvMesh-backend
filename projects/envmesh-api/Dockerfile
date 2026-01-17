FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./
COPY internal ./

RUN go build -o main ./cmd/api

FROM alpine:latest

WORKDIR /

COPY --from=builder /app/main .

COPY docker-entrypoint.sh ./
RUN chmod +x ./docker-entrypoint.sh

EXPOSE 8080

ENTRYPOINT [ "./docker-entrypoint.sh" ]