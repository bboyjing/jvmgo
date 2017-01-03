package classfile

import "encoding/binary"

// byte数组存储读取的字节流
type ClassReader struct {
	data []byte
}

// 读取u1类型的数据
func (self *ClassReader) readUint8() uint8 {
	// 读取第一个字节，8位uint
	val := self.data[0]
	// 将读取过的字节从字节流中剔除
	self.data = self.data[1:]
	return val
}

// 读取u2类型的数据
func (self *ClassReader) readUint16() uint16 {
	//Go标准库提供的解码多字节数据
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

// 读取u4类型的数据
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

// 读取u8类型数据
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

// 读取u2表，表的大小由开头的u2数据指出
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16();
	}
	return s
}

// 读取指定数量的字节
func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}

