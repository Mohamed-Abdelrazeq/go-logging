FROM golang:1.23.2-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go version
RUN go env

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]