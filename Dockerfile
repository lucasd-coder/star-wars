FROM golang:1.18-alpine as build

WORKDIR /go/src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main -ldflags="-s -w -a" ./cmd/app

CMD ["tail","-f","/dev/null"]
