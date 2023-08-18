package web

import (
	"net/http"
	"time"

	"emperror.dev/errors"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

func loggerMiddleware() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			start := time.Now()

			next.ServeHTTP(w, r)

			fields := map[string]interface{}{
				"remoteIP":  r.RemoteAddr,
				"url":       r.URL.Path,
				"proto":     r.Proto,
				"method":    r.Method,
				"userAgent": r.Header.Get("User-Agent"),
				"latency":   time.Since(start).String(),
			}

			if bytesIn := r.Header.Get("Content-Length"); bytesIn != "" {
				fields["bytesIn"] = bytesIn
			}

			if traceID := uuid.UUID(trace.SpanContextFromContext(ctx).TraceID()); traceID != uuid.Nil {
				fields["traceID"] = traceID.String()
			}

			zerolog.Ctx(ctx).Err(ctx.Err()).Str("type", "access").Timestamp().Fields(fields).Send()
		})
	}
}

func recoverMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if re := recover(); re != nil {
					err := errors.Errorf("%s", re)

					zerolog.Ctx(r.Context()).Error().Err(err).Str("type", "panic").Timestamp().Send()

					_, err = w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
					if err != nil {
						zerolog.Ctx(r.Context()).Error().Err(err).Str("type", "panic").Timestamp().Send()
					}

					w.WriteHeader(http.StatusInternalServerError)

					return
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
