# ------------------- BUILD STAGE ------------------- #
FROM golang:1.19-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server

# ------------------- SERVE STAGE ------------------- #

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/config.json .
COPY --from=builder /app/.env .
COPY --from=builder /app/scripts/docker-entry.sh .

RUN chmod +x docker-entry.sh

EXPOSE 3000

CMD ["sh", "docker-entry.sh"]
