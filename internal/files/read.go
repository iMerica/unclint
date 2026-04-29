package files

import (
	"os"
)

type FileContent struct {
	Path    string
	Content string
}

func Read(path string) (*FileContent, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return &FileContent{
		Path:    path,
		Content: string(data),
	}, nil
}
