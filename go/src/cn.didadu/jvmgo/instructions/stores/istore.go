package stores

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// istore结构体
type ISTORE struct{ base.Index8Instruction }

// 将操作数存入指定局部变量表
func (self *ISTORE) Execute(frame *rtdata.Frame) {
	_istore(frame, uint(self.Index))
}

// 统一的istore函数
func _istore(frame *rtdata.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}

// 将操作数存入第0号局部变量表，索引隐含在操作码中
type ISTORE_0 struct{ base.NoOperandsInstruction }

func (self *ISTORE_0) Execute(frame *rtdata.Frame) {
	_istore(frame, 0)
}

// 将操作数存入第1号局部变量表，索引隐含在操作码中
type ISTORE_1 struct{ base.NoOperandsInstruction }

func (self *ISTORE_1) Execute(frame *rtdata.Frame) {
	_istore(frame, 1)
}

// 将操作数存入第2号局部变量表，索引隐含在操作码中
type ISTORE_2 struct{ base.NoOperandsInstruction }

func (self *ISTORE_2) Execute(frame *rtdata.Frame) {
	_istore(frame, 2)
}

// 将操作数存入第3号局部变量表，索引隐含在操作码中
type ISTORE_3 struct{ base.NoOperandsInstruction }

func (self *ISTORE_3) Execute(frame *rtdata.Frame) {
	_istore(frame, 3)
}


