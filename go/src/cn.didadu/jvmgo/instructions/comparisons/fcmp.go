package comparisons

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// fcmpg指令结构体
type FCMPG struct{ base.NoOperandsInstruction }

func (self *FCMPG) Execute(frame *rtdata.Frame) {
	_fcmp(frame, true)
}

// fcmpl指令结构体
type FCMPL struct{ base.NoOperandsInstruction }

func (self *FCMPL) Execute(frame *rtdata.Frame) {
	_fcmp(frame, false)
}

func _fcmp(frame *rtdata.Frame, gFlag bool) {
	stack := frame.OperandStack()
	// 弹出栈顶float型变量，作为比较符后面的操作数
	v2 := stack.PopFloat()
	// 弹出栈顶long型变量，作为比较府前面的操作数
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		// fcmpg指令对于第四种结果入栈1
		stack.PushInt(1)
	} else {
		// fcmpl指令对于第四种结果入栈0
		stack.PushInt(-1)
	}
}