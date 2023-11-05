package cacher

type CacherInterface interface {
	GetById(id string) ([]byte, error)
	Save(id string, data []byte)
}
