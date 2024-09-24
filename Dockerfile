FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/cmd/main .

EXPOSE 80

CMD ["/app/cmd/main"]