# ------------------- BUILD STAGE ------------------- #
FROM golang:1.19-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server

# ------------------- DEV ------------------- #

FROM builder AS dev

WORKDIR /app

RUN apk add --no-cache make

RUN go install github.com/cespare/reflex@latest

ENTRYPOINT ["./scripts/docker-entry.sh"]
CMD ["make watch"]

# ------------------- PROD ------------------- #

FROM alpine:latest as prod
WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/config.json .
COPY --from=builder /app/.env .
COPY --from=builder /app/scripts/docker-entry.sh .

ENTRYPOINT ["./scripts/docker-entry.sh"]
CMD ["./server"]
