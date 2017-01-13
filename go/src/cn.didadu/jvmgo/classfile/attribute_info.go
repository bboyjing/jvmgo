package classfile

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/

type AttributeInfo interface {
	// 读取属性信息由具体的属性结构实现
	readInfo(reader *ClassReader)
}
