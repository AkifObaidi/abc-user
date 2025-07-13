FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o abc-user-backend .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/abc-user-backend /usr/local/bin/abc-user-backend

WORKDIR /usr/local/bin

EXPOSE 8080

CMD ["./abc-user-backend"]
