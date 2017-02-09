package classfile

import "fmt"

type ConstantPool []ConstantInfo

// 读取常量池
func readConstantPool(reader *ClassReader) ConstantPool {
	// 常量池数量(2个字节)
	cpCount := int(reader.readUint16())
	// 生成长度为cpCount，类型是ConstantInfo的slice，就是常量池
	cp := make([]ConstantInfo, cpCount)
	// 注意：常量池的数量从1开始！
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		// ConstantLongInfo和ConstantDoubleInfo占两个位置
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
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

// 通过索引读取ClassName
func (self ConstantPool) getClassName(index uint16) string {
	// 通过索引读取CONSTANT_Class_info
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	// 通过CONSTANT_Class_info的nameIndex读取类的字面值
	return self.getUtf8(classInfo.nameIndex)
}

// 读取ConstantNameAndTypeInfo中name_index和descriptor_index对应的字面值
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	// 通过索引读取ConstantNameAndTypeInfo
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	// 通过ConstantNameAndTypeInfo中nameIndex索引读取字面值
	name := self.getUtf8(ntInfo.nameIndex)
	// 通过ConstantNameAndTypeInfo中descriptorIndex索引读取字面值
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}