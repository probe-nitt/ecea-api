# ------------------- BUILD STAGE ------------------- #

FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init

RUN go build -o server

# ------------------- DEV ------------------- #

FROM builder AS dev

WORKDIR /app

RUN apk add --no-cache make

RUN go install github.com/cespare/reflex@latest

ENV DB_PORT=5000

ENV DB_HOST=ecea_db

ENTRYPOINT ["./scripts/entry.sh"]

CMD ["make watch"]

# ------------------- PROD ------------------- #

FROM alpine:latest AS prod

WORKDIR /app

RUN mkdir static

COPY --from=builder /app/server /app/scripts/entry.sh /app/.env  ./

ENV DB_PORT=5000

ENV DB_HOST=ecea_db

ENTRYPOINT ["/app/scripts/entry.sh"]

CMD [ "./server" ]
