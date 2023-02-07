FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o /go-stripe-api ./cmd/api/*.go

EXPOSE 8083

CMD ["/go-stripe-api"]
