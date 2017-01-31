package base

// BytecodeReader结构体
type BytecodeReader struct {
	// 存放字节码
	code []byte
	// 记录存取到了哪个字节码
	pc   int
}

// 为了避免每次解码指令都新创建一个BytecodeReader实例，所以定义一个Reset()方法
func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

// 获取pc
func (self *BytecodeReader) PC() int {
	return self.pc
}

// 读取一个字节的uint8
func (self *BytecodeReader) ReadUint8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}

// 将uint8转换成int8
func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

// 读取两个字节的uint16
func (self *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(self.ReadUint8())
	byte2 := uint16(self.ReadUint8())
	return (byte1 << 8) | byte2
}

// 将uint16转换成int16
func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}
