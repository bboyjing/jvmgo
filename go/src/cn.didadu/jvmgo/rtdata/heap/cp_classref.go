package heap

import "cn.didadu/jvmgo/classfile"

// 类符号引用结构体
type ClassRef struct {
	SymRef
}

// 初始化类符号引用结构体
func newClassRef(cp *ConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp
	ref.className = classInfo.Name()
	return ref
}
