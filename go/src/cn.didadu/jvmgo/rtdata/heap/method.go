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
		methods[i] = newMethod(class, cfMethod)
	}
	return methods
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	// 分解方法描述符
	md := parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(md.parameterTypes)
	// 如果是本地方法，注入字节码和其他信息
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
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
func (self *Method) calcArgSlotCount(paramTypes []string) {
	for _, paramType := range paramTypes {
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

// 注入字节码等信息
func (self *Method) injectCodeAttribute(returnType string) {
	// 暂定操作数栈深度
	self.maxStack = 4 // todo
	self.maxLocals = self.argSlotCount
	switch returnType[0] {
	case 'V':
		self.code = []byte{0xfe, 0xb1} // return
	case 'L', '[':
		self.code = []byte{0xfe, 0xb0} // areturn
	case 'D':
		self.code = []byte{0xfe, 0xaf} // dreturn
	case 'F':
		self.code = []byte{0xfe, 0xae} // freturn
	case 'J':
		self.code = []byte{0xfe, 0xad} // lreturn
	default:
		self.code = []byte{0xfe, 0xac} // ireturn
	}
}