FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o /go-stripe-api ./cmd/api/*.go

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /go-stripe-api /go-stripe-api

EXPOSE 8083

USER nonroot:nonroot

ENTRYPOINT ["/go-stripe-api"]
