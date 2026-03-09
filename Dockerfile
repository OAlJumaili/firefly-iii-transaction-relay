FROM golang:1.19
WORKDIR /app

COPY go.sum go.mod ./
RUN go mod download

COPY . .
RUN RUN CGO_ENABLED=0 GOOS=linux go build -o ./firefly-relay

EXPOSE 8080
CMD [/app/firefly-relay]