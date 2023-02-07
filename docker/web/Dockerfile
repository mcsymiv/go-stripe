FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o /go-stripe-web ./cmd/web/*.go

EXPOSE 8082

CMD ["/go-stripe-web"]
