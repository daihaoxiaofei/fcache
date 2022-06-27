package storage

import (
	"errors"
	"github.com/boltdb/bolt"
	"path"
)

type BoltStore struct {
	dirPath string
	bucket  []byte
	key     []byte
	db      *bolt.DB
}

var _ Storage = (*BoltStore)(nil)

// NewBoltStore 这个目前还有问题无法使用
// 且也bolt设计理念不符 效率不高
func NewBoltStore(dirPath string) *BoltStore {
	boltStore := &BoltStore{
		dirPath: dirPath,
		bucket:  []byte("bucket"),
		key:     []byte("key"),
	}
	return boltStore
}

func (f *BoltStore) Remove(fileName string) error {
	return f.Write([]byte(``))
}

func (f *BoltStore) Open(fileName string) (err error) {
	f.db, err = bolt.Open(path.Join(f.dirPath, fileName), 0644, nil)
	return err
}

func (f *BoltStore) Close() {
	f.db.Close()
}

func (f *BoltStore) Read() (val []byte, err error) {
	err = f.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(f.bucket)
		if bucket == nil {
			return errors.New("Bucket not found!" + string(f.bucket))
		}
		val = bucket.Get(f.key)
		return nil
	})
	return
}

func (f *BoltStore) Write(result []byte) error {
	return f.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(f.bucket)
		// bucket := tx.Bucket(world)
		if err != nil {
			return err
		}

		err = bucket.Put(f.key, result)
		if err != nil {
			return err
		}
		return nil
	})

}
