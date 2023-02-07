FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o /go-stripe-web ./cmd/web/*.go

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /go-stripe-web /go-stripe-web

EXPOSE 8083

USER nonroot:nonroot

ENTRYPOINT ["/go-stripe-web"]
