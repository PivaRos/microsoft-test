FROM golang:1.22-alpine AS builder
ENV GO111MODULE=on
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./src .
COPY config.yaml .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .
COPY --from=builder /app/config.yaml .
EXPOSE 80
CMD ["./app"]
