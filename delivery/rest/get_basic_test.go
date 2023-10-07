package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/mhdiiilham/resume/delivery/rest"
	"github.com/mhdiiilham/resume/delivery/rest/mock"
	"github.com/mhdiiilham/resume/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type resumeHandlerTestSuite struct {
	suite.Suite
}

func TestResumeHandler(t *testing.T) {
	suite.Run(t, new(resumeHandlerTestSuite))
}

func (suite *resumeHandlerTestSuite) TestGetBasic() {
	testCases := []struct {
		name                string
		expectedMockReturn0 entity.Basic
		expectedErr         error
		expectedResp        map[string]any
	}{
		{
			name: "success",
			expectedMockReturn0: entity.Basic{
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
			expectedErr: nil,
			expectedResp: map[string]any{
				"code":    float64(http.StatusOK),
				"message": http.StatusText(http.StatusOK),
				"data": map[string]any{
					"name":    "Muhammad Ilham",
					"label":   "Backend Engineer",
					"image":   "",
					"email":   "muhd.iiilham@gmail.com",
					"phone":   "(+62) 896-5887-6167",
					"url":     "",
					"summary": "Experienced backend engineer with a passion for programming and a strong track record in the technology industry. Fluent in Go, with experience using PHP, Laravel, and NodeJS to develop scalable, secure backend systems. Skilled in working with Docker, Google Cloud Platform, and AWS, and experienced in implementing continuous integration and deployment strategies using Bitbucket CI/CD and GitHub Actions. Strong logic skills, a love for exploring new technologies, and a commitment to writing high-quality code with comprehensive unit testing and a TDD approach.",
					"location": map[string]any{
						"address":     "",
						"postalCode":  "",
						"city":        "South Jakarta",
						"countryCode": "ID",
						"region":      "DKI Jakarta",
					},
					"profiles": []any{
						map[string]any{
							"network":  "Twitter",
							"username": "w8rloO",
							"url":      "https://twitter.com/w8rloO",
						},
					},
				},
				"error": nil,
			},
		},
	}

	for _, tt := range testCases {
		suite.T().Run(tt.name, func(t *testing.T) {
			assertion := assert.New(t)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockResumeService := mock.NewMockResumeService(ctrl)
			mockResumeService.EXPECT().GetBasic().Return(tt.expectedMockReturn0).Times(1)

			app := fiber.New()
			router := app.Group("/api/v1")
			rest.RegisterFiberHandlers(router, mockResumeService)

			req := httptest.NewRequest(http.MethodGet, "/api/v1/resumes/basics", nil)
			resp, err := app.Test(req)
			assertion.Equal(tt.expectedErr, err)

			var actualResp map[string]any
			err = json.NewDecoder(resp.Body).Decode(&actualResp)
			assertion.NoError(err)

			assertion.Equal(tt.expectedResp, actualResp)
		})
	}

}
