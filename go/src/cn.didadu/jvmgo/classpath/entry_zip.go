package classpath

import (
	"path/filepath"
	"archive/zip"
	"io/ioutil"
	"errors"
)

// ZipEntry表示ZIP或JAR文件形式的路径
type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry  {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	/*
		defer表示延迟执行，此处表示readClass方法return的时候关闭资源
		类似Java中的finally
	 */
	defer r.Close()

	/*
		for循环的range格式可以对slice进行迭代循环
		如果能找到给定路径的文件，则返回，否则返回错误信息
	 */
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}

			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}