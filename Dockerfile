FROM dhi.io/golang:1.26-alpine3.23-sfw-ent-dev
WORKDIR /app

COPY go.sum go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./firefly-relay

EXPOSE 8080
CMD ["/app/firefly-relay"]