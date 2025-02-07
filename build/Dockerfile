FROM golang:alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/gostart

FROM alpine

WORKDIR /app
RUN apk add --no-cache curl
COPY --from=builder /app/main .

CMD ["./main"]
