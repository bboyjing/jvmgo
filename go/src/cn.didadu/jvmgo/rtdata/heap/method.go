package heap

import "cn.didadu/jvmgo/classfile"

type Method struct {
	// 字段和方法公用结构体
	ClassMember
	// 操作数栈最大深度
	maxStack     uint
	// 局部变量表大小
	maxLocals    uint
	// 字节码
	code         []byte
	// 方法的参数所占的Slot数量
	argSlotCount uint
}

// 初始化方法信息
func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
		methods[i].calcArgSlotCount()
	}
	return methods
}

// 从方法的属性中读取maxStack、maxLocals和字节码
func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
	}
}

// Getter方法
func (self *Method) MaxStack() uint {
	return self.maxStack
}
func (self *Method) MaxLocals() uint {
	return self.maxLocals
}
func (self *Method) Code() []byte {
	return self.code
}

// 计算方法参数占用的slot数量
func (self *Method) calcArgSlotCount() {
	// 分解方法描述符
	parsedDescriptor := parseMethodDescriptor(self.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
		self.argSlotCount++
		// long和double占两个slot
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
	}
	if !self.IsStatic() {
		// 如果不是静态方法，给隐藏的参数this添加一个slot位置
		self.argSlotCount++
	}
}

// 获取方法的参数所占的Slot数量
func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}

func (self *Method) IsAbstract() bool {
	return 0 != self.accessFlags & ACC_ABSTRACT
}

func (self *Method) IsNative() bool {
	return 0 != self.accessFlags & ACC_NATIVE
}