# BUILD
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o library-go

# RUNNER
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/library-go .

EXPOSE 9000
CMD [ "./library-go" ]