package grpc

import (
	"context"
	"encoding/binary"
	"path"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func RecoverUnaryServerInterceptor(ctx context.Context, enableStack bool) grpc.UnaryServerInterceptor {
	var customFunc grpc_recovery.RecoveryHandlerFunc = func(p interface{}) error {
		e := status.Errorf(codes.Internal, "[PANIC] %s\n\n%s", p, string(debug.Stack()))

		zerolog.Ctx(ctx).
			Error().
			Err(e).
			Str("type", "panic").
			Timestamp().
			Send()

		if enableStack {
			return e
		}

		return status.Errorf(codes.Internal, "Internal Server Error")
	}

	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(customFunc),
	}

	return grpc_recovery.UnaryServerInterceptor(opts...)
}

func AccessLogUnaryServerInterceptor(mainCtx context.Context) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()

		resp, err := handler(ctx, req)

		fields := map[string]interface{}{
			"remote_ip":    getRemoteIP(ctx),
			"method":       getMethodName(info.FullMethod),
			"user_agent":   getUserAgent(ctx),
			"status_error": getStatusError(err),
			"latency":      time.Since(start).String(),
			"bytes_in":     getMessageBytesCount(req),
			"bytes_out":    getMessageBytesCount(resp),
		}

		if traceID := uuid.UUID(trace.SpanContextFromContext(ctx).TraceID()); traceID != uuid.Nil {
			fields["trace_id"] = traceID.String()
		}

		zerolog.Ctx(mainCtx).Err(ctx.Err()).Str("type", "access").Timestamp().Fields(fields).Send()

		return resp, err
	}
}

func getRemoteIP(ctx context.Context) string {
	if p, ok := peer.FromContext(ctx); ok {
		return p.Addr.String()
	}

	return ""
}

func getMethodName(method string) string {
	return path.Base(method)
}

func getUserAgent(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		return strings.Join(md.Get("user-agent"), "")
	}

	return ""
}

func getStatusError(err error) string {
	return status.Convert(err).String()
}

func getMessageBytesCount(message interface{}) string {
	if pb, ok := message.(proto.Message); ok {
		if b, err := protojson.Marshal(pb); err == nil {
			return strconv.Itoa(binary.Size(b))
		}
	}

	return ""
}
