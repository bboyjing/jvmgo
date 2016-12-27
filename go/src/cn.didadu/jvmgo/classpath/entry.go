package classpath

import (
	"os"
	"strings"
)

/*
	获取系统分隔符
	Windows下是;类UNIX系统下是:
 */
const pathListSeparator = string(os.PathListSeparator)

// 定义Entry接口
type Entry interface {
	// 负责寻找和加载class文件
	readClass(className string) ([]byte, Entry, error)
	// 返回变量的字符串表示，相当于Java中的toString
	String() string
}

func newEntry(path string) Entry {
	//如果path中包含分隔符，则实例化CompositeEntry
	if (strings.Contains(path, pathListSeparator)) {
		return newCompositeEntry(path)
	}
	//如果path以*结尾，则实例化WildcardEntry
	if (strings.HasSuffix(path, "*")) {
		return newWildcardEntry(path)
	}
	//如果path以jar或者zip结尾，则返回ZipEntry
	if strings.HasSuffix(path, ".jar") ||
		strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") ||
		strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	//以上都不匹配则返回DirEntry
	return newDirEntry(path)
}