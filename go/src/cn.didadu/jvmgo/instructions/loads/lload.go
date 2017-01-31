package loads

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)


// lload指令结构体
type LLOAD struct{ base.Index8Instruction }

// 通过索操作局部变量表
func (self *LLOAD) Execute(frame *rtdata.Frame) {
	_lload(frame, uint(self.Index))
}

// 统一的lload函数
func _lload(frame *rtdata.Frame, index uint) {
	// 通过索引读取局部变量表
	val := frame.LocalVars().GetLong(index)
	// 将局部变量表中的值推入栈顶
	frame.OperandStack().PushLong(val)
}

// 操作第0号局部变量，索引隐含在操作码中
type LLOAD_0 struct{ base.NoOperandsInstruction }
func (self *LLOAD_0) Execute(frame *rtdata.Frame) {
	_lload(frame, 0)
}

// 操作第1号局部变量，索引隐含在操作码中
type LLOAD_1 struct{ base.NoOperandsInstruction }
func (self *LLOAD_1) Execute(frame *rtdata.Frame) {
	_lload(frame, 1)
}

// 操作第2号局部变量，索引隐含在操作码中
type LLOAD_2 struct{ base.NoOperandsInstruction }
func (self *LLOAD_2) Execute(frame *rtdata.Frame) {
	_lload(frame, 2)
}

// 操作第3号局部变量，索引隐含在操作码中
type LLOAD_3 struct{ base.NoOperandsInstruction }
func (self *LLOAD_3) Execute(frame *rtdata.Frame) {
	_lload(frame, 3)
}


