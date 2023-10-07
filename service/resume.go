package service

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mhdiiilham/resume/entity"
)

type ResumeService struct {
	fileReader FileReader
	filename   string
	resume     entity.Resume
}

func NewResumeService(fileReader FileReader, filename string) *ResumeService {
	return &ResumeService{
		fileReader: fileReader,
		filename:   filename,
	}
}

func (s *ResumeService) loadResume() error {
	bytesResume, err := s.fileReader.Read(s.filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytesResume, &s.resume)
}

func (s *ResumeService) GetBasic() (entity.Basic, error) {
	if err := s.loadResume(); err != nil {
		return entity.Basic{}, err
	}

	return s.resume.Basics, nil
}
func (s *ResumeService) GetWhatsAppDotMeURL() (string, error) {
	if err := s.loadResume(); err != nil {
		return "", err
	}

	phoneNumber := s.resume.Basics.Phone
	phoneNumber = strings.ReplaceAll(phoneNumber, "(", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, ")", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "+", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")
	return fmt.Sprintf("https://wa.me/%s", phoneNumber), nil
}
