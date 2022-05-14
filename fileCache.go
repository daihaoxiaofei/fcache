package fcache

import (
	"errors"
	"fcache/coder"
	"io/ioutil"
	"os"
	"path"
)

const (
	defaultPath = `cache`
)

// expanded-name
type FileCache struct {
	code    coder.Coder
	DirPath string
	Suffix  string // 后缀名
	// mu      sync.Mutex // 多线程调用时加锁会影响效率 和 意外情况全局锁死(如嵌套缓存) 所以得靠调用者自己控制了
}
type RememberFunc func() interface{}

// 为了调用方便
var DefaultFC = NewFC(defaultPath)
var Remember = DefaultFC.Remember

func NewFC(dirPath string) *FileCache {
	return &FileCache{
		code:    coder.NewJsonCode(),
		DirPath: dirPath,
		Suffix:  `.db`,
	}
}

// 清除缓存
func (f *FileCache) Remome(key string) {
	_ = os.Remove(path.Join(f.DirPath, key+f.Suffix))
}

func (f *FileCache) Remember(key string, out interface{}, fun RememberFunc) (err error) {
	if _, err := os.Stat(f.DirPath); err != nil {
		err := os.MkdirAll(f.DirPath, os.ModePerm)
		if err != nil {
			panic(f.DirPath + ` 路径创建失败 Mkdir err: ` + err.Error())
		}
	}
	file, err := os.OpenFile(path.Join(f.DirPath, key+f.Suffix), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return
	}
	defer file.Close()

	result, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	if len(result) == 0 {
		result, err = f.code.EnCode(fun())
		if err != nil {
			return
		}
		if result == nil {
			file.Close() // 关闭文件
			_ = os.Remove(path.Join(f.DirPath, key+f.Suffix)) // 删除文件
			return errors.New(`函数返回为空`)
		}
		_, err = file.Write(result)
		if err != nil {
			return
		}
	}
	f.code.DeCode(result, out)

	return
}
