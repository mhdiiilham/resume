package rest

import "github.com/gofiber/fiber/v2"

func RegisterFiberHandlers(r fiber.Router, srv ResumeService) {
	resumeRouter := r.Group("/resumes")
	resumeRouter.Get("/basics", getBasicHandler(srv))
}
