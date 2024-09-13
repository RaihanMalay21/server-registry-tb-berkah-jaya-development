
# Stage 1: Build
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o server_registration main.go

# Stage 2: Runtime
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server_registration .

RUN chmod +x /app/server_registration

USER nobody
CMD ["/app/server_registration"]

EXPOSE 8080


