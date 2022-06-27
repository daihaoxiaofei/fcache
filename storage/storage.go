package storage

type Storage interface {
	Open(string) error
	Close()
	Remove(string) error
	Read() ([]byte, error)
	Write([]byte) error
}
