package references

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// arraylength指令结构体
type ARRAY_LENGTH struct{ base.NoOperandsInstruction }

func (self *ARRAY_LENGTH) Execute(frame *rtdata.Frame) {
	// 获取操作数栈，并且弹出栈顶引用
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	// 若引用为空，抛出NullPointerException异常
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}
	// 获取数组长度
	arrLen := arrRef.ArrayLength()
	// 讲数组长度推入操作数栈顶
	stack.PushInt(arrLen)
}
