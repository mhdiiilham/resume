package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func getBasicHandler(srv ResumeService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		basic := srv.GetBasic()
		return ctx.JSON(HTTPResponse{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    basic,
			Error:   nil,
		})
	}
}
