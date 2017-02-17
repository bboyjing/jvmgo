package constants

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// ldc指令结构体(单字节操作数)
type LDC struct{ base.Index8Instruction }

func (self *LDC) Execute(frame *rtdata.Frame) {
	_ldc(frame, self.Index)
}

func _ldc(frame *rtdata.Frame, index uint) {
	// 获取操作数栈
	stack := frame.OperandStack()
	// 获取运行时常量池
	cp := frame.Method().Class().ConstantPool()
	// 通过索引从常量池中获取常量值
	c := cp.GetConstant(index)

	// 根据常量值的类型将对应的值入栈
	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	// case string:
	// case *heap.ClassRef:
	// case MethodType, MethodHandle
	default:
		panic("todo: ldc!")
	}
}

// ldc_w指令结构体(两个字节操作数)
type LDC_W struct{ base.Index16Instruction }

func (self *LDC_W) Execute(frame *rtdata.Frame) {
	_ldc(frame, self.Index)
}

// ldc2_w指令结构体(两个字节操作数)
type LDC2_W struct{ base.Index16Instruction }

func (self *LDC2_W) Execute(frame *rtdata.Frame) {
	// 获取操作数栈
	stack := frame.OperandStack()
	// 获取运行时常量池
	cp := frame.Method().Class().ConstantPool()
	// 通过索引从常量池中获取常量值
	c := cp.GetConstant(self.Index)

	// 根据常量值的类型将对应的值入栈
	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}