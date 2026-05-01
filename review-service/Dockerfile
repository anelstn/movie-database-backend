FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/bin/review-service ./review-service/main.go

FROM alpine:3.22

WORKDIR /app
RUN apk add --no-cache tzdata
ENV TZ=Asia/Almaty
COPY --from=builder /app/bin/review-service /app/review-service

EXPOSE 8081

CMD ["/app/review-service"]
