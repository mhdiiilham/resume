package service_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/mhdiiilham/resume/entity"
	"github.com/mhdiiilham/resume/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type resumeServiceTestSuite struct {
	suite.Suite
}

func TestResumeService(t *testing.T) {
	suite.Run(t, new(resumeServiceTestSuite))
}

func (suite *resumeServiceTestSuite) TestGetBasic() {
	testCases := []struct {
		name                        string
		resume                      entity.Resume
		expectedMockArg0            string
		generateExpectedMockReturn0 func(resume entity.Resume) ([]byte, error)
		expectedMockReturn1         error
		expectedErr                 error
	}{
		{
			name: "success",
			resume: entity.Resume{
				Basics: entity.Basic{
					Name:    "Muhammad Ilham",
					Label:   "Backend Engineer",
					Image:   "",
					Email:   "muhd.iiilham@gmail.com",
					Phone:   "(+62) 896-5887-6167",
					Summary: "Experienced backend engineer with a passion for programming and a strong track record in the technology industry. Fluent in Go, with experience using PHP, Laravel, and NodeJS to develop scalable, secure backend systems. Skilled in working with Docker, Google Cloud Platform, and AWS, and experienced in implementing continuous integration and deployment strategies using Bitbucket CI/CD and GitHub Actions. Strong logic skills, a love for exploring new technologies, and a commitment to writing high-quality code with comprehensive unit testing and a TDD approach.",
					Location: entity.Location{
						Address:     "",
						PostalCode:  "",
						City:        "South Jakarta",
						CountryCode: "ID",
						Region:      "DKI Jakarta",
					},
					Profiles: []entity.Profile{
						{
							Network:  "Twitter",
							Username: "w8rloO",
							URL:      "https://twitter.com/w8rloO",
						},
					},
				},
			},
			expectedMockArg0: "resume.json",
			generateExpectedMockReturn0: func(resume entity.Resume) ([]byte, error) {
				return json.Marshal(resume)
			},
			expectedMockReturn1: nil,
			expectedErr:         nil,
		},
		{
			name:             "file not found",
			resume:           entity.Resume{},
			expectedMockArg0: "resume.json",
			generateExpectedMockReturn0: func(resume entity.Resume) ([]byte, error) {
				return nil, os.ErrNotExist
			},
			expectedMockReturn1: os.ErrExist,
			expectedErr:         os.ErrExist,
		},
	}

	for _, tt := range testCases {
		suite.T().Run(tt.name, func(t *testing.T) {
			assertion := assert.New(t)

			s := service.NewResumeService(tt.resume)
			assertion.Equal(tt.resume.Basics, s.GetBasic())
		})
	}
}

func (suite *resumeServiceTestSuite) TestGetWhatsAppURL() {
	testCases := []struct {
		name                string
		resume              entity.Resume
		expectedWhatsAppURL string
	}{
		{
			name: "success",
			resume: entity.Resume{
				Basics: entity.Basic{
					Phone: "(+62) 896-5887-6167",
				},
			},
			expectedWhatsAppURL: "https://wa.me/6289658876167",
		},
	}

	for _, tt := range testCases {
		suite.T().Run(tt.name, func(t *testing.T) {
			assertion := assert.New(t)
			s := service.NewResumeService(tt.resume)
			assertion.Equal(tt.expectedWhatsAppURL, s.GetWhatsAppDotMeURL())
		})
	}
}
