package lang

import (
	"cn.didadu/jvmgo/native"
	"cn.didadu/jvmgo/rtdata"
	"cn.didadu/jvmgo/rtdata/heap"
)

func init() {
	native.Register("java/lang/System", "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
}

// public static native void arraycopy(Object src, int srcPos, Object dest, int destPos, int length)
// (Ljava/lang/Object;ILjava/lang/Object;II)V
func arraycopy(frame *rtdata.Frame) {
	// 获取局部变量表的5个参数
	vars := frame.LocalVars()
	src := vars.GetRef(0)
	srcPos := vars.GetInt(1)
	dest := vars.GetRef(2)
	destPos := vars.GetInt(3)
	length := vars.GetInt(4)

	// 源数组和目标数组都不能为空
	if src == nil || dest == nil {
		panic("java.lang.NullPointerException")
	}

	// 检查源数组和目标数组是否兼容
	if !checkArrayCopy(src, dest) {
		panic("java.lang.ArrayStoreException")
	}

	// 验证参数是否合法
	if srcPos < 0 || destPos < 0 || length < 0 ||
		srcPos+length > src.ArrayLength() ||
		destPos+length > dest.ArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}

	// 拷贝数组
	heap.ArrayCopy(src, dest, srcPos, destPos, length)
}

// 判断源数组和目标数组是否兼容
func checkArrayCopy(src, dest *heap.Object) bool {
	srcClass := src.Class()
	destClass := dest.Class()

	// 源数组和目标数组都必须是数组
	if !srcClass.IsArray() || !destClass.IsArray() {
		return false
	}

	// 如果数组是基本类型，则必须是相同类型
	if srcClass.ComponentClass().IsPrimitive() ||
		destClass.ComponentClass().IsPrimitive() {
		return srcClass == destClass
	}
	return true
}