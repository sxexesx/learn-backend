package storage

import (
	"github.com/google/uuid"
	"github.com/sxexesx/learn-backend/go/_projects/08_modules/server/internal/file"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(fileName string) (*file.File, error) {
	newFile, err := file.NewFile(fileName)
	if err != nil {
		return nil, err
	}

	s.files[newFile.ID] = newFile

	return newFile, nil

}
