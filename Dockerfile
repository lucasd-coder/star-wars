FROM golang:1.18-alpine as build
WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main -ldflags="-s -w -a" ./cmd/app

FROM alpine:latest as compressed
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache upx

COPY --from=build /src /src
WORKDIR /src
RUN upx main

# Expose port
EXPOSE ${HTTP_PORT}

FROM scratch
COPY --from=compressed /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=compressed /src /
ENTRYPOINT [ "/main" ]
