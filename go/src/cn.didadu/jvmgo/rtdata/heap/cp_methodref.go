package heap

import "cn.didadu/jvmgo/classfile"

// 方法引用结构体
type MethodRef struct {
	MemberRef
	method *Method
}

// 初始化方法符号引用
func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

