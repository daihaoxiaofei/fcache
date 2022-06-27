package fcache

import (
	"github.com/daihaoxiaofei/fcache/coder"
	"github.com/daihaoxiaofei/fcache/storage"
	"reflect"
)

const (
	defaultPath = `cache`
)

// FileCache expanded-name
type FileCache struct {
	coder   coder.Coder
	Storage storage.Storage
	Suffix  string // 后缀名
	// mu      sync.Mutex // 多线程调用时加锁会影响效率 和 意外情况全局锁死(如嵌套缓存) 所以得靠调用者自己控制了
}
type RememberFunc func() interface{}

var DefaultFC = NewFC(defaultPath)
var Remember = DefaultFC.Remember
var Remove = DefaultFC.Remove

func NewFC(dirPath string) *FileCache {
	return &FileCache{
		coder:   coder.NewGobCode(),
		Storage: storage.NewFileStore(dirPath),
		Suffix:  `.db`,
	}
}

// Remove 清除缓存
func (f *FileCache) Remove(key string) {
	_ = f.Storage.Remove(key + f.Suffix)
}

func (f *FileCache) Remember(key string, out interface{}, fun RememberFunc) (err error) {
	err = f.Storage.Open(key + f.Suffix)
	if err != nil {
		return
	}
	defer f.Storage.Close()

	result, err := f.Storage.Read()
	if err != nil {
		return
	}

	if len(result) == 0 {
		res := fun()
		// 将res赋值给out
		reflect.ValueOf(out).Elem().Set(reflect.ValueOf(res))

		result, err = f.coder.EnCode(res)
		if err != nil {
			return
		}
		err = f.Storage.Write(result)
		return
	}
	return f.coder.DeCode(result, out)
}
