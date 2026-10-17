FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN apk add --no-cache git
RUN go mod download
COPY . .
RUN go build -o /usr/local/bin/ember-gate ./cmd/server

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/local/bin/ember-gate /usr/local/bin/ember-gate
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/ember-gate"]
