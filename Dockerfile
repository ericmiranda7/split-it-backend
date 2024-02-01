FROM golang:latest as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .
# Static build; no C libraries required.
RUN CGO_ENABLED=0 GOOS=linux go build -v

FROM alpine:latest

COPY --from=builder /app/split-that /split-that

ENTRYPOINT ["/split-that"]