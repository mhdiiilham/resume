package service

import (
	"encoding/json"

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

func (s *ResumeService) LoadResume() error {
	bytesResume, err := s.fileReader.Read(s.filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytesResume, &s.resume)
}

func (s *ResumeService) GetBasic() entity.Basic {
	return s.resume.Basics
}
