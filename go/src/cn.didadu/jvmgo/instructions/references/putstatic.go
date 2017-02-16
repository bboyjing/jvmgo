package references

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
	"cn.didadu/jvmgo/rtdata/heap"
)

//putstatic 指令结构体(2个字节操作数)
type PUT_STATIC struct{ base.Index16Instruction }

func (self *PUT_STATIC) Execute(frame *rtdata.Frame) {
	// 获取当前正在执行的方法
	currentMethod := frame.Method()
	// 获取当前类
	currentClass := currentMethod.Class()
	// 获取运行时常量池
	cp := currentClass.ConstantPool()
	// 通过索引从常量池中获取字段符号引用
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	// 解析字段
	field := fieldRef.ResolvedField()
	class := field.Class()

	// 判断字段是否是static
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 判断字段是否是final
	if field.IsFinal() {
		// 如果字段是final，则表示静态常量，只能在类初始化方法中赋值
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	// 获取字段描述符,也就是字段的类型
	descriptor := field.Descriptor()
	// 获取字段在Slots中的索引
	slotId := field.SlotId()
	// 获取类的静态变量Slots
	slots := class.StaticVars()
	// 获取操作数栈
	stack := frame.OperandStack()

	// 通过字段的类型从操作数栈顶弹出相应的值，并给Class的静态变量Slots赋值
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	default:
	}
}
