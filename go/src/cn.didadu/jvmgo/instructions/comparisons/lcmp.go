package comparisons

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// long比较指令结构体
type LCMP struct{ base.NoOperandsInstruction }

func (self *LCMP) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	// 弹出栈顶long型变量，作为比较符后面的操作数
	v2 := stack.PopLong()
	// 弹出栈顶long型变量，作为比较府前面的操作数
	v1 := stack.PopLong()
	if v1 > v2 {
		// 若v1大于v2，则将1入栈
		stack.PushInt(1)
	} else if v1 == v2 {
		// 若v1等于v2，则将0入栈
		stack.PushInt(0)
	} else {
		// 若v1小于v2，则将-1入栈
		stack.PushInt(-1)
	}
}