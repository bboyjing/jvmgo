package classpath

import (
	"path/filepath"
	"io/ioutil"
)

// 定义DirEntry结构体
type DirEntry struct {
	// 用于存放绝对路径
	absDir string
}

/*
	返回指向DirEntry对象的指针
	Go语言没有专门的构造函数，此函数就当做DirEntry的构造函数
 */
func newDirEntry(path string) *DirEntry {
	/*
		Go使用error值来表示错误状态
		Go使用panic和recover来处理所谓的运行时异常
	 */
	absDir, err := filepath.Abs(path)
	if (err != nil) {
		panic(err)
	}
	return &DirEntry{absDir}
}

/*
	Go没有类，可以使用方法接受者的方式在结构体类型上定义方法
	指向DirEntry对象的指针self为方法接受者
	该方法用来读取class文件文件
 */
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

// 实现String方法
func (self *DirEntry) String() string {
	return self.absDir
}