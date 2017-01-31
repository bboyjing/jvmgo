package base

import "cn.didadu/jvmgo/rtdata"

// 定义指令接口
type Instruction interface {
	// 从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)
	// 执行指令
	Execute(frame *rtdata.Frame)
}

/*
	表示没有操作数的指令
	没有任何字段，FetchOperands()方法也为空
 */
type NoOperandsInstruction struct {
}
func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
}

// 表示跳转指令，Offset字段存储跳转偏移量
type BranchInstruction struct {
	// 存储跳转偏移量
	Offset int
}
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	// 从字节码中读取uint16整数，转成int后赋给Offset
	self.Offset = int(reader.ReadInt16())
}

// 存储和加载指令需要根据索引存取局部变量表，索引由单字节操作数给出
type Index8Instruction struct {
	//局部变量表索引
	Index uint
}
func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	// 从字节码中读取一个uint8整数，转成uint后赋给Index
	self.Index = uint(reader.ReadUint8())
}

// 有一些指令需要访问运行时常量池，常量池索引由2个字节操作数给出
type Index16Instruction struct {
	// 常量池索引
	Index uint
}
func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	// 读取一个uint16整数，转成uint后赋给Index
	self.Index = uint(reader.ReadUint16())
}

