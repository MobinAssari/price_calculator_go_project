package iomanager

type IOManager interface {
	WriteResult(data any) error
	ReadFile() ([]string, error)
}
