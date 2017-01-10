package classfile

/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type ClassFile struct {
	magic      uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

// 读取魔数
func (self *ClassFile) readAndCheckMagic(reader *ClassReader)  {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		/*
			Java虚拟机的实现是抛出java.lang.ClassFormatError
			目前由于才开始写虚拟机，还没法抛出异常，暂先调用panic
		 */
		panic("java.lang.ClassFormatError: magic!")
	}
}

// 读取版本号
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch  self.majorVersion {
	// 主版本号为45直接返回
	case 45:
		return
	// 主版本号为46~52时，次版本号必须为0
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// 读取类访问标志
func (self *ClassFile) readAccessFlags(reader *ClassReader) {
	self.accessFlags = reader.readUint16()
}

// 读取类和超类索引
func (self *ClassFile) readThisClass(reader *ClassReader) {
	self.thisClass = reader.readUint16()
}

func (self *ClassFile) readSuperClass(reader *ClassReader) {
	self.superClass = reader.readUint16()
}

// 读取接口索引表
func (self *ClassFile) readInterface(reader *ClassReader) {
	//接口是u2类型的表结构，所以用的readUint16s方法读取
	self.interfaces = reader.readUint16s()
}

// 读取字段表
func (self *ClassFile) readFields(reader *ClassReader) {
	readMembers(reader, self.constantPool)
}

// 读取方法表
func (self *ClassFile) readMethods(reader *ClassReader) {
	readMembers(reader, self.constantPool)
}