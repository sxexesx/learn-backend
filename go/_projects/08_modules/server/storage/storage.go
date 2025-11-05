package storage

type Storage struct {
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Upload(fileName string, fileData []byte) (*file.File, error) {

}
