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

// 跳过Padding字节
func (self *BytecodeReader) SkipPadding() {
	for self.pc%4 != 0 {
		self.ReadUint8()
	}
}

// 读取指定数量的int32，并返回数组
func (self *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = self.ReadInt32()
	}
	return ints
}

// 读取4个字节uint8转成int32
func (self *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(self.ReadUint8())
	byte2 := int32(self.ReadUint8())
	byte3 := int32(self.ReadUint8())
	byte4 := int32(self.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}