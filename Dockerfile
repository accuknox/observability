FROM golang:1.17 AS builder
WORKDIR /home/observability
COPY . .
RUN go mod tidy -compat=1.17
# RUN CGO_ENABLED=0 GOOS=linux go build -o main .
RUN apk add build-base
RUN apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init
RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o main main.go
FROM artifactory.accuknox.com/accuknox/ubuntu:22.04
WORKDIR /home/observability
RUN mkdir config
COPY config/app.yaml config/app.yaml
COPY --from=builder /home/observability/main /home/observability/main
CMD ["./main"]
