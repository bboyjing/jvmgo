package heap

// 判断是否是数组类
func (self *Class) IsArray() bool {
	return self.name[0] == '['
}

// 创建数组
func (self *Class) NewArray(count uint) *Object {
	if !self.IsArray() {
		panic("Not array class: " + self.name)
	}
	switch self.Name() {
	case "[Z":
		return &Object{self, make([]int8, count), nil}
	case "[B":
		return &Object{self, make([]int8, count), nil}
	case "[C":
		return &Object{self, make([]uint16, count), nil}
	case "[S":
		return &Object{self, make([]int16, count), nil}
	case "[I":
		return &Object{self, make([]int32, count), nil}
	case "[J":
		return &Object{self, make([]int64, count), nil}
	case "[F":
		return &Object{self, make([]float32, count), nil}
	case "[D":
		return &Object{self, make([]float64, count), nil}
	default:
		return &Object{self, make([]*Object, count), nil}
	}
}

// 获取数组元素类名，并加载类
func (self *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(self.name)
	return self.loader.LoadClass(componentClassName)
}

func ArrayCopy(src, dst *Object, srcPos, dstPos, length int32) {
	switch src.data.(type) {
	case []int8:
		_src := src.data.([]int8)[srcPos : srcPos + length]
		_dst := dst.data.([]int8)[dstPos : dstPos + length]
		copy(_dst, _src)
	case []int16:
		_src := src.data.([]int16)[srcPos : srcPos + length]
		_dst := dst.data.([]int16)[dstPos : dstPos + length]
		copy(_dst, _src)
	case []int32:
		_src := src.data.([]int32)[srcPos : srcPos + length]
		_dst := dst.data.([]int32)[dstPos : dstPos + length]
		copy(_dst, _src)
	case []int64:
		_src := src.data.([]int64)[srcPos : srcPos + length]
		_dst := dst.data.([]int64)[dstPos : dstPos + length]
		copy(_dst, _src)
	case []uint16:
		_src := src.data.([]uint16)[srcPos : srcPos + length]
		_dst := dst.data.([]uint16)[dstPos : dstPos + length]
		copy(_dst, _src)
	case []float32:
		_src := src.data.([]float32)[srcPos : srcPos + length]
		_dst := dst.data.([]float32)[dstPos : dstPos + length]
		copy(_dst, _src)
	case []float64:
		_src := src.data.([]float64)[srcPos : srcPos + length]
		_dst := dst.data.([]float64)[dstPos : dstPos + length]
		copy(_dst, _src)
	case []*Object:
		_src := src.data.([]*Object)[srcPos : srcPos + length]
		_dst := dst.data.([]*Object)[dstPos : dstPos + length]
		copy(_dst, _src)
	default:
		panic("Not array!")
	}
}