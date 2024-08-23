package MiddleWare

import (
	logger "github.com/SHERATONS/OMS-Sellsuki-Internship/Observability/Log"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"time"
)

func LoggerMiddleWare(c *fiber.Ctx) error {
	start := time.Now()

	fields := logrus.Fields{
		"method":   c.Method(),
		"path":     c.Path(),
		"query":    c.OriginalURL(),
		"remoteIP": c.IP(),
	}

	if err := c.Next(); err != nil {
		err = c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		if err != nil {
			return err
		}
		return err
	}

	fields["status"] = c.Response().StatusCode()
	fields["latency"] = time.Since(start).Seconds()

	logger.LogInfo("HTTP request", fields)

	return nil
}
