FROM artifactory.accuknox.com/accuknox/golang:1.16.5-alpine3.14 AS builder
WORKDIR /home/observability
COPY . .
RUN go mod tidy
RUN GOOS=linux go build -o main .

FROM artifactory.accuknox.com/accuknox/ubuntu:22.04
WORKDIR /home/observability
COPY --from=builder /home/observability/main /home/observability/main
CMD ["./main"]