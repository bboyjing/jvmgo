package constants

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// Push byte
type BIPUSH struct {
	val int8
}

// 读取单个字节操作数
func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

// 将操作数int值推入栈顶
func (self *BIPUSH) Execute(frame *rtdata.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}


// Push short
type SIPUSH struct {
	val int16
}

// 读取两个字节操作数
func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}
// 将操作数int值推入栈顶
func (self *SIPUSH) Execute(frame *rtdata.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}