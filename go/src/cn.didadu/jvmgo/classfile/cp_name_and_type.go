package classfile

/*
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
*/


type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

// 读取CONSTANT_NameAndType_info
func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	// 读取2个字节的name_index
	self.nameIndex = reader.readUint16()
	// 读取2个字节的descriptorIndex
	self.descriptorIndex = reader.readUint16()
}