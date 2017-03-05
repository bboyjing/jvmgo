package stores

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// astore结构体
type ASTORE struct{ base.Index8Instruction }

// 统一的astore函数
func _astore(frame *rtdata.Frame, index uint) {
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}

// 将操作数存入指定局部变量表
func (self *ASTORE) Execute(frame *rtdata.Frame) {
	_astore(frame, uint(self.Index))
}

// 将操作数存入第0号局部变量表，索引隐含在操作码中
type ASTORE_0 struct{ base.NoOperandsInstruction }

func (self *ASTORE_0) Execute(frame *rtdata.Frame) {
	_astore(frame, 0)
}

// 将操作数存入第1号局部变量表，索引隐含在操作码中
type ASTORE_1 struct{ base.NoOperandsInstruction }

func (self *ASTORE_1) Execute(frame *rtdata.Frame) {
	_astore(frame, 1)
}

// 将操作数存入第2号局部变量表，索引隐含在操作码中
type ASTORE_2 struct{ base.NoOperandsInstruction }

func (self *ASTORE_2) Execute(frame *rtdata.Frame) {
	_astore(frame, 2)
}

// 将操作数存入第3号局部变量表，索引隐含在操作码中
type ASTORE_3 struct{ base.NoOperandsInstruction }

func (self *ASTORE_3) Execute(frame *rtdata.Frame) {
	_astore(frame, 3)
}