package logger

import (
	"context"
	"log/slog"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var log *slog.Logger

const request_id = "request_id"

type Logger interface {
	Debug(log string, args ...any)
	DebugWithCtx(ctx context.Context, logMsg string, args ...any)
	Info(log string, args ...any)
	InfoWithCtx(ctx context.Context, logMsg string, args ...any)
	Error(log string, args ...any)
	ErrorWithCtx(ctx context.Context, logMsg string, args ...any)
	Warn(log string, args ...any)
	WarnWithCtx(ctx context.Context, logMsg string, args ...any)
	AttachRequestIdToRequests(c *gin.Context)
}

type Impl struct {
}

/*
Debug log
*/
func Debug(logMsg string, args ...any) {
	log.Debug(logMsg, args...)
}

func DebugWithCtx(ctx context.Context, logMsg string, args ...any) {
	args = append(args, request_id, ctx.Value(request_id))
	log.DebugContext(ctx, logMsg, args...)
}

/*
Info log
*/
func Info(logMsg string, args ...any) {
	log.Info(logMsg, args...)
}

func InfoWithCtx(ctx context.Context, logMsg string, args ...any) {
	args = append(args, request_id, ctx.Value(request_id))
	log.InfoContext(ctx, logMsg, args...)
}

/*
Warn log
*/
func Warn(logMsg string, args ...any) {
	log.Warn(logMsg, args...)
}

func WarnWithCtx(ctx context.Context, logMsg string, args ...any) {
	args = append(args, request_id, ctx.Value(request_id))
	log.WarnContext(ctx, logMsg, args...)
}

/*
Error log
*/
func Error(logMsg string, args ...any) {
	log.Error(logMsg, args...)
}

func ErrorWithCtx(ctx context.Context, logMsg string, args ...any) {
	args = append(args, request_id, ctx.Value(request_id))
	log.ErrorContext(ctx, logMsg, args...)
}

func AttachRequestIdToRequests(c *gin.Context) {
	u := uuid.New()
	c.Set(request_id, u)
	DebugWithCtx(c, "request started", "method", c.Request.Method, "path", c.Request.URL.Path)
	c.Next()
}

func Init(logLevel string) {
	handlerOptions := slog.HandlerOptions{Level: getLogLevel(logLevel)}
	log = slog.New(slog.NewJSONHandler(os.Stdout, &handlerOptions))
	slog.SetDefault(log)
}

func getLogLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
