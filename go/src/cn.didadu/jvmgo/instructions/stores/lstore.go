package stores

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// lstore结构体
type LSTORE struct{ base.Index8Instruction }

// 将操作数存入指定局部变量表
func (self *LSTORE) Execute(frame *rtdata.Frame) {
	_lstore(frame, uint(self.Index))
}

// 统一的lstore函数
func _lstore(frame *rtdata.Frame, index uint) {
	// 弹出位于操作数栈顶的操作数
	val := frame.OperandStack().PopLong()
	// 设置局部变量
	frame.LocalVars().SetLong(index, val)
}

// 将操作数存入第0号局部变量表，索引隐含在操作码中
type LSTORE_0 struct{ base.NoOperandsInstruction }

func (self *LSTORE_0) Execute(frame *rtdata.Frame) {
	_lstore(frame, 0)
}

// 将操作数存入第1号局部变量表，索引隐含在操作码中
type LSTORE_1 struct{ base.NoOperandsInstruction }

func (self *LSTORE_1) Execute(frame *rtdata.Frame) {
	_lstore(frame, 1)
}

// 将操作数存入第2号局部变量表，索引隐含在操作码中
type LSTORE_2 struct{ base.NoOperandsInstruction }

func (self *LSTORE_2) Execute(frame *rtdata.Frame) {
	_lstore(frame, 2)
}

// 将操作数存入第3号局部变量表，索引隐含在操作码中
type LSTORE_3 struct{ base.NoOperandsInstruction }

func (self *LSTORE_3) Execute(frame *rtdata.Frame) {
	_lstore(frame, 3)
}


