package service

import (
	"fmt"
	"regexp"

	"github.com/mhdiiilham/resume/entity"
)

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(`[^0-9]`)
}

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
	phoneNumber := re.ReplaceAllString(s.resume.Basics.Phone, "")
	return fmt.Sprintf("https://wa.me/%s", phoneNumber)
}
