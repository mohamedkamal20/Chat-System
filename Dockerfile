FROM golang:1.18-alpine

WORKDIR /Chat-System

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
