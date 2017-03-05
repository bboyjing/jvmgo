package rtdata

import "math"
import "cn.didadu/jvmgo/rtdata/heap"

// 局部变量表结构体
type LocalVars []Slot

// 创建LocalVars实例
func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

/*
	设置、读取int
 */
func (self LocalVars) SetInt(index uint, val int32) {
	self[index].num = val
}
func (self LocalVars) GetInt(index uint) int32 {
	return self[index].num
}

/*
	设置、读取float
 */
func (self LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}
func (self LocalVars) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

/*
	设置、读取long
	需要两个slot
 */
func (self LocalVars) SetLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index + 1].num = int32(val >> 32)
}
func (self LocalVars) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index + 1].num)
	return int64(high) << 32 | int64(low)
}

/*
	设置、读取double
	需要两个slot
 */
func (self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}
func (self LocalVars) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

// 设置、读取引用
func (self LocalVars) SetRef(index uint, ref *heap.Object) {
	self[index].ref = ref
}
func (self LocalVars) GetRef(index uint) *heap.Object {
	return self[index].ref
}

// 设置
func (self LocalVars) SetSlot(index uint, slot Slot) {
	self[index] = slot
}