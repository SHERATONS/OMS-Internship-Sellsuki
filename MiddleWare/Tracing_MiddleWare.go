package MiddleWare

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func TracingMiddleWare(c *fiber.Ctx) error {
	propagator := otel.GetTextMapPropagator()
	ctx := propagator.Extract(context.Background(), c.Request().Header)
	tracer := otel.Tracer("github.com/SHERATONS/OMS-Sellsuki-Internship")

	ctx, span := tracer.Start(ctx, "HTTP "+c.Method()+" "+c.Path())
	defer span.End()

	// Add request attributes
	span.SetAttributes(
		attribute.String("http.method", c.Method()),
		attribute.String("http.url", c.OriginalURL()),
		attribute.String("http.user_agent", c.Get("User-Agent")),
		attribute.String("http.client_ip", c.IP()),
	)

	// Pass context with span to the next middleware/handler
	c.SetUserContext(ctx)
	err := c.Next()

	// Add response attributes
	span.SetAttributes(
		attribute.Int("http.status_code", c.Response().StatusCode()),
	)

	return err
	//ctx, span := tracer.Start(c.Context(), "HTTP "+c.Method()+" "+c.Path())
	//defer span.End()
	//
	//span.SetAttributes(
	//	attribute.String("http.method", c.Method()),
	//	attribute.String("http.url", c.OriginalURL()),
	//	attribute.String("http.user_agent", c.Get("user_agent")),
	//	attribute.String("http.client_ip", c.IP()),
	//)
	//
	//c.SetUserContext(ctx)
	//err := c.Next()
	//
	//span.SetAttributes(
	//	attribute.Int("http.status_code", c.Response().StatusCode()),
	//)
	//
	//return err
}
