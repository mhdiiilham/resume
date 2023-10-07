package service

import (
	"fmt"
	"strings"

	"github.com/mhdiiilham/resume/entity"
)

type ResumeService struct {
	resume entity.Resume
}

func NewResumeService(resume entity.Resume) *ResumeService {
	return &ResumeService{
		resume: resume,
	}
}

func (s *ResumeService) GetBasic() entity.Basic {
	return s.resume.Basics
}

func (s *ResumeService) GetWhatsAppDotMeURL() string {
	phoneNumber := s.resume.Basics.Phone
	phoneNumber = strings.ReplaceAll(phoneNumber, "(", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, ")", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "+", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")
	return fmt.Sprintf("https://wa.me/%s", phoneNumber)
}
