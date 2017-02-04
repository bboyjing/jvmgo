package math

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/rtdata"
)

// innc结构体
type IINC struct {
	// 常量池索引
	Index uint
	// 常量值
	Const int32
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	// 读取单字节操作数，作为局部变量表索引
	self.Index = uint(reader.ReadUint8())
	// 读取单字节操作数，作为常量值
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtdata.Frame) {
	localVars := frame.LocalVars()
	// 通过索引获取局部变量
	val := localVars.GetInt(self.Index)
	// 局部变量 + 常量
	val += self.Const
	// 回写局部变量
	localVars.SetInt(self.Index, val)
}

