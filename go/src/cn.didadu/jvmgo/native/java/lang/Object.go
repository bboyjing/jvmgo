package lang

import (
	"cn.didadu/jvmgo/native"
	"cn.didadu/jvmgo/rtdata"
	"unsafe"
)

const jlObject = "java/lang/Object"

func init() {
	native.Register(jlObject, "getClass", "()Ljava/lang/Class;", getClass)
	native.Register(jlObject, "hashCode", "()I", hashCode)
	native.Register(jlObject, "clone", "()Ljava/lang/Object;", clone)
}

/*
	public final native Class<?> getClass();
	()Ljava/lang/Class;
 */
func getClass(frame *rtdata.Frame) {
	// 从局部变量表拿到this引用
	this := frame.LocalVars().GetThis()
	// 获取类对象
	class := this.Class().JClass()
	// 将类对象入栈
	frame.OperandStack().PushRef(class)
}

// public native int hashCode();
// ()I
func hashCode(frame *rtdata.Frame) {
	// 获取局部变量表第一个元素，this引用
	this := frame.LocalVars().GetThis()
	// 将对象引用转换成int32
	hash := int32(uintptr(unsafe.Pointer(this)))
	// 将计算后的hash code入栈
	frame.OperandStack().PushInt(hash)
}

// protected native Object clone() throws CloneNotSupportedException;
// ()Ljava/lang/Object;
func clone(frame *rtdata.Frame) {
	// 获取局部变量表第一个元素，this引用
	this := frame.LocalVars().GetThis()

	// 加载Cloneable类
	cloneable := this.Class().Loader().LoadClass("java/lang/Cloneable")
	if !this.Class().IsImplements(cloneable) {
		panic("java.lang.CloneNotSupportedException")
	}

	// 将克隆后的对象引用入栈
	frame.OperandStack().PushRef(this.Clone())
}