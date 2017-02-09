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
