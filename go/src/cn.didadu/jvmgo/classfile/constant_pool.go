package classfile

import "fmt"

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	return nil
}

// 通过索引读取常量结构体
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	/*
		为什么不是index - 1？
		因为常量池的索引是从1开始的！
	 */
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic(fmt.Errorf("Invalid constant pool index: %v!", index))
}

// 通过索引获取ConstantUtf8Info结构体的字面值
func (self ConstantPool)getUtf8(index uint16) string {
	// 将ConstantInfo转换为实际类型ConstantUtf8Info
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}