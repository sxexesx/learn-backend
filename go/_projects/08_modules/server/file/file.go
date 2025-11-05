package file

import (
	"fmt"
	"github.com/google/uuid"
)

type File struct {
	ID   uuid.UUID
	Name string
}

func NewFile(name string) (*File, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &File{
		ID:   id,
		Name: name,
	}, nil
}

func (f *File) String() string {
	return fmt.Sprintf("File (%s, %v)", f.Name, f.ID)
}
