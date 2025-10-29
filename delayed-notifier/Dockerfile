FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" go build -o service ./cmd/

FROM alpine:3.22 AS runner

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/service ./service

USER 1000:1000

CMD ["./service"]