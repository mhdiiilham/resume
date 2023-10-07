package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/mhdiiilham/resume/delivery/rest"
	"github.com/mhdiiilham/resume/entity"
	"github.com/mhdiiilham/resume/service"
)

const PORT = ":8081"

func main() {
	var filename string

	flag.StringVar(&filename, "file", "resume.json", "load resume from .json file")
	flag.Parse()

	resume, err := loadResume(filename)
	if err != nil {
		log.Fatalf("failed to open resume file json: %s; error=%s", filename, err.Error())
	}

	// init services here
	resumeService := service.NewResumeService(resume)

	app := fiber.New()
	router := app.Group("/api/v1")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Println("Gracefully shut down")
		err = app.Shutdown()
		log.Printf("shutting down app; err=%v", err)
	}()

	// /api/v1/resumes
	rest.RegisterFiberHandlers(router, resumeService)

	if err := app.Listen(PORT); err != nil {
		log.Fatalf("failed to listen app at PORT=%s; err=%v", PORT, err)
	}

	log.Println("running cleanup tasks...")
}

func loadResume(filename string) (entity.Resume, error) {
	var resume entity.Resume

	resumeBytes, err := os.ReadFile(filename)
	if err != nil {
		return entity.Resume{}, err
	}

	if err := json.Unmarshal(resumeBytes, &resume); err != nil {
		return entity.Resume{}, err
	}

	return resume, nil
}
