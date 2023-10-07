package service

//go:generate mockgen -destination=mock/file_reader_mock.go -package=mock github.com/mhdiiilham/resume/service FileReader
type FileReader interface {
	Read(path string) ([]byte, error)
}
