package references

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
	"cn.didadu/jvmgo/rtdata/heap"
)

// putfield指令结构体(2个字节操作数)
type PUT_FIELD struct{ base.Index16Instruction }

func (self *PUT_FIELD) Execute(frame *rtdata.Frame) {
	// 获取当前正在执行的方法
	currentMethod := frame.Method()
	// 获取当前类
	currentClass := currentMethod.Class()
	// 获取运行时常量池
	cp := currentClass.ConstantPool()
	// 通过索引从常量池中获取字段符号引用
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()

	// 判断字段是否是static
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 判断字段是否是final
	if field.IsFinal() {
		// 如果字段是final，则表示静态常量，只能在类初始化方法中赋值
		if currentClass != field.Class() || currentMethod.Name() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	// 获取字段描述符,也就是字段的类型
	descriptor := field.Descriptor()
	// 获取字段在Slots中的索引
	slotId := field.SlotId()
	// 获取操作数栈
	stack := frame.OperandStack()

	/*
		根据字段类型从操作数栈弹出相应的变量，然后弹出对象引用
		判断引用是否是null，若是抛出NullPointerException异常，否则通过引用实例变量赋值
	 */
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		val := stack.PopInt()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetInt(slotId, val)
	case 'F':
		val := stack.PopFloat()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetFloat(slotId, val)
	case 'J':
		val := stack.PopLong()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetLong(slotId, val)
	case 'D':
		val := stack.PopDouble()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetDouble(slotId, val)
	case 'L', '[':
		val := stack.PopRef()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetRef(slotId, val)
	default:
	}
}