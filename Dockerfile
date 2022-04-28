FROM artifactory.accuknox.com/accuknox/golang:1.16.5-alpine3.14 AS builder
WORKDIR /home/observability
COPY . .
RUN go mod tidy
# RUN CGO_ENABLED=0 GOOS=linux go build -o main .
RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o main main.go
FROM artifactory.accuknox.com/accuknox/ubuntu:22.04
WORKDIR /home/observability
RUN mkdir config
COPY config/app.yaml config/app.yaml
COPY --from=builder /home/observability/main /home/observability/main
CMD ["./main"]
