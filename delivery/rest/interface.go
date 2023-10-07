package rest

import "github.com/mhdiiilham/resume/entity"

//go:generate mockgen -destination=mock/resume_service_mock.go -package=mock . ResumeService
type ResumeService interface {
	GetBasic() entity.Basic
}
