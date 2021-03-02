FROM golang:1.14-alpine
WORKDIR /merchant

RUN apk update && apk add --no-cache gcc musl-dev git mysql-client

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app

COPY deployment/docker/bin /usr/local/bin/merchant/
RUN chmod +x /usr/local/bin/merchant/* && chmod +x /merchant/bin/*

CMD ["/merchant/bin/app"]
EXPOSE 8080