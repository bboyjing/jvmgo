package classfile

/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
*/

type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

// 读取SourceFile_attribute
func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	// attribute_length的值为2，所以固定读取2个字节
	self.sourceFileIndex = reader.readUint16()
}

// 通过索引读取常量池中utf8字面值
func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}