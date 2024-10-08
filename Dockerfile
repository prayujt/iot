FROM golang:1.22.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o iot .


FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/iot .

CMD ["./iot"]

