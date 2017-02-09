package heap

import "cn.didadu/jvmgo/classfile"

// 字段信息结构体
type Field struct {
	// 字段和方法公用结构体
	ClassMember
	// 字段在slot中的位置
	slotId          uint
	// 常量表达式的值的索引
	constValueIndex uint
}

// 初始化字段信息
func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

// 判断访问标识
func (self *Field) IsVolatile() bool {
	return 0 != self.accessFlags&ACC_VOLATILE
}
func (self *Field) IsTransient() bool {
	return 0 != self.accessFlags&ACC_TRANSIENT
}
func (self *Field) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

// slotId Getter
func (self *Field) SlotId() uint {
	return self.slotId
}

// 判断字段类型是否是long或double
func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

// constValueIndex Getter
func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}

// 从class文件中复制field_info的属性，此处只读取constValueIndex
func (self *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		self.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}