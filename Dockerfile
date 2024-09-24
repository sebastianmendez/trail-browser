FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/cmd/main /app/cmd

EXPOSE 80

CMD ["/app/cmd/main"]