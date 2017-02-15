package heap

import "cn.didadu/jvmgo/classfile"

// 字段和方法公用结构体
type ClassMember struct {
	// 访问标志
	accessFlags uint16
	// 名字
	name        string
	// 描述符
	descriptor  string
	// 当前类指针，可以通过字段或方法访问到它所属的类
	class       *Class
}

// 从class文件中复制数据
func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

// 判断访问标识
func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}
func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}
func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}
func (self *ClassMember) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *ClassMember) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}

// Getter方法
func (self *ClassMember) Name() string {
	return self.name
}
func (self *ClassMember) Descriptor() string {
	return self.descriptor
}
func (self *ClassMember) Class() *Class {
	return self.class
}

// 判断self字段是否能被d class访问
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	// 如果self字段是public，则任何类都可以访问
	if self.IsPublic() {
		return true
	}
	// 获取self字段所属的class
	c := self.class
	// 如果self字段是protected，只有子类和同一包下的类可以访问
	if self.IsProtected() {
		return d == c || d.isSubClassOf(c) ||
			c.getPackageName() == d.getPackageName()
	}
	// 如果是默认访问权限，则只有同一包下的类可以访问
	if !self.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	// self字段是private，则只有本类才可以访问
	return d == c
}