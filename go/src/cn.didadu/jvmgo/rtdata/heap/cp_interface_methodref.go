package heap

import "cn.didadu/jvmgo/classfile"

// 接口方法符号引用结构体
type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

// 初始化接口方法符号引用
func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}