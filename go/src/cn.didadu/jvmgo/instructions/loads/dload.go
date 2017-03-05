package loads

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// dload指令结构体
type DLOAD struct{ base.Index8Instruction }

// 通过索操作局部变量表
func (self *DLOAD) Execute(frame *rtdata.Frame) {
	_dload(frame, uint(self.Index))
}

// 统一的dload函数
func _dload(frame *rtdata.Frame, index uint) {
	// 通过索引读取局部变量表
	val := frame.LocalVars().GetDouble(index)
	// 将局部变量表中的值推入栈顶
	frame.OperandStack().PushDouble(val)
}

// 操作第0号局部变量，索引隐含在操作码中
type DLOAD_0 struct{ base.NoOperandsInstruction }

func (self *DLOAD_0) Execute(frame *rtdata.Frame) {
	_dload(frame, 0)
}

// 操作第1号局部变量，索引隐含在操作码中
type DLOAD_1 struct{ base.NoOperandsInstruction }

func (self *DLOAD_1) Execute(frame *rtdata.Frame) {
	_dload(frame, 1)
}

// 操作第2号局部变量，索引隐含在操作码中
type DLOAD_2 struct{ base.NoOperandsInstruction }

func (self *DLOAD_2) Execute(frame *rtdata.Frame) {
	_dload(frame, 2)
}

// 操作第3号局部变量，索引隐含在操作码中
type DLOAD_3 struct{ base.NoOperandsInstruction }

func (self *DLOAD_3) Execute(frame *rtdata.Frame) {
	_dload(frame, 3)
}