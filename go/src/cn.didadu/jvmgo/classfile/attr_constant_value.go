package classfile

/*
ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 constantvalue_index;
}
*/

type ConstantValueAttribute struct {
	constantValueIndex uint16
}

// 读取ConstantValue_attribute
func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	// 读取读取2个字节的constantValueIndex
	self.constantValueIndex = reader.readUint16()
}

// 读取constantValueIndex
func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}