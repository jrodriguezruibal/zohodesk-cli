# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o zohodesk-cli .

# Final stage
FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /root/

COPY --from=builder /app/zohodesk-cli /usr/local/bin/zohodesk-cli

RUN mkdir -p /root/.config/zohodesk-cli

ENTRYPOINT ["zohodesk-cli"]
CMD ["version"]