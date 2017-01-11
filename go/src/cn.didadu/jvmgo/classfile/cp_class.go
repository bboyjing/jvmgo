package classfile

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

// 读取CONSTANT_Class_info
func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	//读取name_index(2个字节)
	self.nameIndex = reader.readUint16()
}

// 读取class字面值
func (self *ConstantClassInfo) Name() string {
	// 需要通过nameIndex到常量池读取utf8字面值
	return self.cp.getUtf8(self.nameIndex)
}