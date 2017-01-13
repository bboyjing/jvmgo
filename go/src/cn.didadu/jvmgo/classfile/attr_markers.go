package classfile

// 统一的结构体
type MarkerAttribute struct{}

/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type SyntheticAttribute struct {
	MarkerAttribute
}

// attribute_length值为0，所以方法中什么都不做
func (self *MarkerAttribute) readInfo(reader *ClassReader) {
}
