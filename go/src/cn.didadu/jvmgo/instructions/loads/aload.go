package loads

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)


// aload指令结构体
type ALOAD struct{ base.Index8Instruction }

// 通过索操作局部变量表
func (self *ALOAD) Execute(frame *rtdata.Frame) {
	_aload(frame, uint(self.Index))
}

// 统一的aload函数
func _aload(frame *rtdata.Frame, index uint) {
	// 通过索引读取局部变量表
	ref := frame.LocalVars().GetRef(index)
	// 将局部变量表中的值推入栈顶
	frame.OperandStack().PushRef(ref)
}

// 操作第0号局部变量，索引隐含在操作码中
type ALOAD_0 struct{ base.NoOperandsInstruction }
func (self *ALOAD_0) Execute(frame *rtdata.Frame) {
	_aload(frame, 0)
}

// 操作第1号局部变量，索引隐含在操作码中
type ALOAD_1 struct{ base.NoOperandsInstruction }
func (self *ALOAD_1) Execute(frame *rtdata.Frame) {
	_aload(frame, 1)
}

// 操作第2号局部变量，索引隐含在操作码中
type ALOAD_2 struct{ base.NoOperandsInstruction }
func (self *ALOAD_2) Execute(frame *rtdata.Frame) {
	_aload(frame, 2)
}

// 操作第3号局部变量，索引隐含在操作码中
type ALOAD_3 struct{ base.NoOperandsInstruction }
func (self *ALOAD_3) Execute(frame *rtdata.Frame) {
	_aload(frame, 3)
}
