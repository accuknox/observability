FROM artifactory.accuknox.com/accuknox/golang:1.16.5-alpine3.14 AS builder
WORKDIR /home/observability
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM artifactory.accuknox.com/accuknox/ubuntu:22.04
WORKDIR /home/observability
RUN apt-get update && apt-get install -y apt-transport-https ca-certificates
COPY --from=builder /home/observability/main /home/observability/main
CMD ["./main"]