package classfile

import "math"

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/

type ConstantIntegerInfo struct {
	val int32
}

// 读取CONSTANT_Integer_info
func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	// 读取4个字节
	bytes := reader.readUint32()
	// 将bytes转为int32
	self.val = int32(bytes)
}

// 获取int32常量值
func (self *ConstantIntegerInfo) Value() int32 {
	return self.val
}


/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
*/

type ConstantFloatInfo struct {
	val float32
}

// 读取CONSTANT_Float_info
func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	// 读取4个字节
	bytes := reader.readUint32()
	// 将bytes转换为float32
	self.val = math.Float32frombits(bytes)
}

// 获取float32常量值
func (self *ConstantFloatInfo) Value() float32 {
	return self.val
}

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/

type ConstantLongInfo struct {
	val int64
}

// 读取CONSTANT_Long_info
func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	// 读取8个字节
	bytes := reader.readUint64()
	// 将bytes转换为int64
	self.val = int64(bytes)
}

// 获取int64常量值
func (self *ConstantLongInfo) Value() int64 {
	return self.val
}

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantDoubleInfo struct {
	val float64
}

// 读取CONSTANT_Double_info
func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	// 读取8个字节
	bytes := reader.readUint64()
	// 将bytes转换为float64
	self.val = math.Float64frombits(bytes)
}

// 获取float64常量值
func (self *ConstantDoubleInfo) Value() float64 {
	return self.val
}

