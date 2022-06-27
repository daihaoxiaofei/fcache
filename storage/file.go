package storage

import (
	"io/ioutil"
	"os"
	"path"
)


type FileStore struct {
	dirPath string
	file    *os.File
}

var _ Storage = (*FileStore)(nil)

func NewFileStore(dirPath string) *FileStore {
	fileStore := &FileStore{
		dirPath: dirPath,
	}
	return fileStore
}

func (f *FileStore) Remove(fileName string) error {
	return os.Remove(path.Join(f.dirPath, fileName))
}
func (f *FileStore) Open(fileName string) error {
	_, err := os.Stat(f.dirPath)
	if err != nil {
		err = os.MkdirAll(f.dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	f.file, err = os.OpenFile(path.Join(f.dirPath, fileName), os.O_RDWR|os.O_CREATE, os.ModePerm)
	return err
}

func (f *FileStore) Close() {
	f.file.Close()
}

func (f *FileStore) Read() ([]byte, error) {
	return ioutil.ReadAll(f.file)
}

func (f *FileStore) Write(result []byte) error {
	_, err := f.file.Write(result)
	return err
}
