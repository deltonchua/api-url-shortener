FROM golang:1.21.4-alpine3.18 AS builder

WORKDIR /app

COPY . .

RUN go build -o server .

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/server /app/

EXPOSE 8080

CMD ["./server"]
