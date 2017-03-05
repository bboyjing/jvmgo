package stores

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)


// fstore结构体
type FSTORE struct{ base.Index8Instruction }

// 将操作数存入指定局部变量表
func (self *FSTORE) Execute(frame *rtdata.Frame) {
	_fstore(frame, uint(self.Index))
}

// 统一的fstore函数
func _fstore(frame *rtdata.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

// 将操作数存入第0号局部变量表，索引隐含在操作码中
type FSTORE_0 struct{ base.NoOperandsInstruction }

func (self *FSTORE_0) Execute(frame *rtdata.Frame) {
	_fstore(frame, 0)
}

// 将操作数存入第1号局部变量表，索引隐含在操作码中
type FSTORE_1 struct{ base.NoOperandsInstruction }

func (self *FSTORE_1) Execute(frame *rtdata.Frame) {
	_fstore(frame, 1)
}

// 将操作数存入第2号局部变量表，索引隐含在操作码中
type FSTORE_2 struct{ base.NoOperandsInstruction }

func (self *FSTORE_2) Execute(frame *rtdata.Frame) {
	_fstore(frame, 2)
}

// 将操作数存入第3号局部变量表，索引隐含在操作码中
type FSTORE_3 struct{ base.NoOperandsInstruction }

func (self *FSTORE_3) Execute(frame *rtdata.Frame) {
	_fstore(frame, 3)
}
