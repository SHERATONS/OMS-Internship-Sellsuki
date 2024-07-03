package MiddleWare

import (
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"strconv"
	"time"
)

func TracingMiddleWare(c *fiber.Ctx) error {
	tracer := otel.Tracer("fiber-server")
	ctx, span := tracer.Start(c.Context(), "HTTP "+c.Method()+" "+c.Path())
	defer span.End()

	startTime := time.Now()

	span.SetAttributes(
		attribute.String("http.method", c.Method()),
		attribute.String("http.url", c.OriginalURL()),
		attribute.String("http.host", c.Hostname()),
		attribute.String("http.path", c.OriginalURL()),
	)

	c.SetUserContext(ctx)
	err := c.Next()

	duration := time.Since(startTime)

	span.SetAttributes(
		attribute.String("http.status_code", strconv.Itoa(c.Response().StatusCode())),
		attribute.String("request.duration_ms", duration.String()),
	)

	return err
}
