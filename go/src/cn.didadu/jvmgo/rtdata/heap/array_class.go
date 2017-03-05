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
		return &Object{self, make([]int8, count)}
	case "[B":
		return &Object{self, make([]int8, count)}
	case "[C":
		return &Object{self, make([]uint16, count)}
	case "[S":
		return &Object{self, make([]int16, count)}
	case "[I":
		return &Object{self, make([]int32, count)}
	case "[J":
		return &Object{self, make([]int64, count)}
	case "[F":
		return &Object{self, make([]float32, count)}
	case "[D":
		return &Object{self, make([]float64, count)}
	default:
		return &Object{self, make([]*Object, count)}
	}
}

// 获取数组元素类名，并加载类
func (self *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(self.name)
	return self.loader.LoadClass(componentClassName)
}