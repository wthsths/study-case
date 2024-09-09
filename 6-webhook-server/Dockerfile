FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/myapp .
EXPOSE 8443
CMD ["./myapp"]
