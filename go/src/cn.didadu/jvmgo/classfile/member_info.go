package classfile

/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
	    attribute_info attributes[attributes_count];
}
*/

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

// 读取字段表或者方法表
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	// 获取fields_count或者methods_count
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	// 读取每一个field或者method
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	// 读取class文件中对应的字节
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		// 读取属性表后面再讲
		attributes:      readAttributes(reader, cp),
	}
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}

// 读取Code属性
func (self *MemberInfo) CodeAttribute() *CodeAttribute {
	// 遍历一个method_info中的attributes
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			// 返回CodeAttribute
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}