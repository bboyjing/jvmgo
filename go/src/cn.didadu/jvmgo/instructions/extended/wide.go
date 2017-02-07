package extended

import (
	"cn.didadu/jvmgo/instructions/base"
	"cn.didadu/jvmgo/instructions/loads"
	"cn.didadu/jvmgo/instructions/stores"
	"cn.didadu/jvmgo/instructions/math"
	"cn.didadu/jvmgo/rtdata"
)

/*
	wide指令结构体
	wide指令只是增加了索引宽度，并不改变子指令操作
 */
type WIDE struct {
	// 存放被改变的指令
	modifiedInstruction base.Instruction
}

func (self *WIDE) FetchOperands(reader *base.BytecodeReader) {
	// 获取操作码
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15:
		// 初始化iload指令结构体(创建子指令实例)
		inst := &loads.ILOAD{}
		// 扩展iload1个字节操作数至两个字节
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x16:
		inst := &loads.LLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x17:
		inst := &loads.FLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x18:
		inst := &loads.DLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x19:
		inst := &loads.ALOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x36:
		inst := &stores.ISTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x37:
		inst := &stores.LSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x38:
		inst := &stores.FSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x39:
		inst := &stores.DSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x3a:
		inst := &stores.ASTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x84:
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		self.modifiedInstruction = inst
	case 0xa9: // ret
		panic("Unsupported opcode: 0xa9!")
	}
}

// 调用子指令的Execute()方法，不改变子指令行为
func (self *WIDE) Execute(frame *rtdata.Frame) {
	self.modifiedInstruction.Execute(frame)
}