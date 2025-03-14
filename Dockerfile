FROM golang:1.24.1-alpine

WORKDIR /app

#COPY go.mod go.sum ./
COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o phonebook-api ./cmd/api

EXPOSE 8080

CMD ["./phonebook-api"]