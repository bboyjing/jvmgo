package references

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
	"cn.didadu/jvmgo/rtdata/heap"
)

// getstatic 指令结构体(2个字节操作数)
type GET_STATIC struct{ base.Index16Instruction }

func (self *GET_STATIC) Execute(frame *rtdata.Frame) {
	// 获取运行时常量池
	cp := frame.Method().Class().ConstantPool()
	// 通过索引从常量池中获取字段符号引用
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	// 解析字段引用
	field := fieldRef.ResolvedField()
	class := field.Class()

	// 判断类初始化是否已经开始
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	// 判断是否是静态方法
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 获取字段描述符,也就是字段的类型
	descriptor := field.Descriptor()
	// 获取字段在Slots中的索引
	slotId := field.SlotId()
	// 获取类的静态变量Slots
	slots := class.StaticVars()
	// 获取操作数栈
	stack := frame.OperandStack()

	// 从静态变量Slots中获取相应的类型，然后推入操作数栈顶
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	default:
	}
}