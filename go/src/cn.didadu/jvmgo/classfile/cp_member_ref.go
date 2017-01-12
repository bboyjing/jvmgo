package classfile



/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
// Go语言没有继承概念，通过结构体嵌套来实现类似继承
type ConstantFieldrefInfo struct{ ConstantMemberrefInfo }

/*
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantMethodrefInfo struct{ ConstantMemberrefInfo }

/*
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantInterfaceMethodrefInfo struct{ ConstantMemberrefInfo }


/*
	CONSTANT_Fieldref_info、CONSTANT_Methodref_info和CONSTANT_InterfaceMethodref_info 的结构一样
	所以使用一样的结构体ConstantMemberrefInfo来表示
 */
type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

// 读取ConstantMemberrefInfo
func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	// 读取2个字节的class_index
	self.classIndex = reader.readUint16()
	// 读取2个字节的name_and_type_index
	self.nameAndTypeIndex = reader.readUint16()
}

// 通过classIndex读取class字面值
func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

// 通过nameAndTypeIndex读取名称和描述字面值
func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}