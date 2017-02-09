package heap

import "cn.didadu/jvmgo/classfile"

// 字段和方法符号引用共用结构体
type MemberRef struct {
	SymRef
	// 字段或方法的名字
	name       string
	// 字段或方法的描述符
	descriptor string
}

// 从class文件中复制数据
func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	// 读取类名
	self.className = refInfo.ClassName()
	// 读取字段或方法的名字和描述符
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}

// Getter方法
func (self *MemberRef) Name() string {
	return self.name
}
func (self *MemberRef) Descriptor() string {
	return self.descriptor
}

