ARG BUILD_ARG_GO_VERSION=1.21
ARG BUILD_ARG_ALPINE_VERSION=3.18
FROM golang:${BUILD_ARG_GO_VERSION}-alpine${BUILD_ARG_ALPINE_VERSION} as builder
WORKDIR /app
COPY . .
RUN apk add --update --no-cache git make \
    && make build-app

FROM alpine:${BUILD_ARG_ALPINE_VERSION}

WORKDIR /app

RUN addgroup -S appgroup && adduser -S appuser -G appgroup \
    && apk --no-cache add ca-certificates \
    && apk add --update --no-cache tzdata

COPY --from=builder /app/bin/app /app/app

RUN chown -R appuser:appgroup /app

USER appuser

# Start the application
CMD ["/app/app"]
