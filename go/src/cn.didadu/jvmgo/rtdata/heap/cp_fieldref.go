package heap

import "cn.didadu/jvmgo/classfile"

// 字段符号引用结构体
type FieldRef struct {
	MemberRef
	// 字段信息结构体
	field *Field
}

// 初始化字段符号引用
func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

// 解析字段符号引用
func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}

func (self *FieldRef) resolveFieldRef() {
	// 获取当前Class指针
	d := self.cp.class
	// 解析字段符号引用之前需要先解析字段所属的类
	c := self.ResolvedClass()
	//根据字段名和描述符查找字段
	field := lookupField(c, self.name, self.descriptor)

	if field == nil {
		// 若没有找到字段，报出NoSuchFieldError异常
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		// 若当前类没有该字段的访问权限，则报出IllegalAccessError异常
		panic("java.lang.IllegalAccessError")
	}

	self.field = field
}

// 根据字段名和描述符查找字段
func lookupField(c *Class, name, descriptor string) *Field {
	// 从Class结构体中遍历fields
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	// 遍历接口
	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}

	// 遍历超类的fields
	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}

	return nil
}