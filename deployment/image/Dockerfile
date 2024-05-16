FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/main cmd/main.go

FROM alpine:3.19  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/bin/main /root/
RUN chmod +x /root/main

CMD ["./main"]