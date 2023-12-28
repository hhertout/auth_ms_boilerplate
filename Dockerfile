FROM golang:1.20-alpine

RUN apk update && apk upgrade
RUN apk add --no-cache  \
    make

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY . .

RUN go mod download

CMD ["make", "watch"]