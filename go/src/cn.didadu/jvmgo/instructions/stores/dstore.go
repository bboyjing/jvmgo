package stores

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// dstore结构体
type DSTORE struct{ base.Index8Instruction }

// 将操作数存入指定局部变量表
func (self *DSTORE) Execute(frame *rtdata.Frame) {
	_dstore(frame, uint(self.Index))
}

// 统一的dstore函数
func _dstore(frame *rtdata.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

// 将操作数存入第0号局部变量表，索引隐含在操作码中
type DSTORE_0 struct{ base.NoOperandsInstruction }
func (self *DSTORE_0) Execute(frame *rtdata.Frame) {
	_dstore(frame, 0)
}

// 将操作数存入第1号局部变量表，索引隐含在操作码中
type DSTORE_1 struct{ base.NoOperandsInstruction }
func (self *DSTORE_1) Execute(frame *rtdata.Frame) {
	_dstore(frame, 1)
}

// 将操作数存入第2号局部变量表，索引隐含在操作码中
type DSTORE_2 struct{ base.NoOperandsInstruction }
func (self *DSTORE_2) Execute(frame *rtdata.Frame) {
	_dstore(frame, 2)
}

// 将操作数存入第3号局部变量表，索引隐含在操作码中
type DSTORE_3 struct{ base.NoOperandsInstruction }
func (self *DSTORE_3) Execute(frame *rtdata.Frame) {
	_dstore(frame, 3)
}