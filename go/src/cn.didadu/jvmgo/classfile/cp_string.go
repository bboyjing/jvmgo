package classfile

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

// 读取CONSTANT_String_info
func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	// 读取string_index（2个字节）
	self.stringIndex = reader.readUint16()
}

// 读取string字面常量
func (self *ConstantStringInfo) String() string{
	//需要通过stringIndex到常量池中读取
	return self.cp.getUtf8(self.stringIndex)
}