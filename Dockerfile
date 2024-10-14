
FROM golang:1.21.3-alpine AS builder
RUN mkdir /build
ADD go.mod go.sum main.go /build/
ADD routes/ /build/routes/
WORKDIR /build
RUN go build

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/prattl-web /app/
COPY public/ /app/public
WORKDIR /app
CMD ["./prattl-web"]

