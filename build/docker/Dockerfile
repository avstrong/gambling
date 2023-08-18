FROM golang:1.20.3-alpine3.16 as builder
WORKDIR /app
COPY . .
RUN apk add --update --no-cache git make
RUN make build-app

FROM alpine:3.17.3

WORKDIR /app

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
RUN apk --no-cache add ca-certificates
RUN apk add --update --no-cache tzdata

COPY --from=builder /app/bin/app /app/app

RUN chown -R appuser:appgroup /app

USER appuser

# Start the application
CMD ["/app/app"]
