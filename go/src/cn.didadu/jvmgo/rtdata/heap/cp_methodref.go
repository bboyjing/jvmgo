package heap

import "cn.didadu/jvmgo/classfile"

// 方法引用结构体
type MethodRef struct {
	MemberRef
	method *Method
}

// 初始化方法符号引用
func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

// 解析非接口方法符号引用
func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

func (self *MethodRef) resolveMethodRef() {
	// 获取当前Class指针
	d := self.cp.class
	// 解析方法符号引用之前需要先解析方法所属的类
	c := self.ResolvedClass()
	// 判断方法所属类是否是接口
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 根据方法名和描述符查找方法
	method := lookupMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	// 判断方法方法权限
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.method = method
}

func lookupMethod(class *Class, name, descriptor string) *Method {
	// 先从方法所属class中查找方法
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		// 若方法所属class中没有找到，再从方法所属class的接口中去找
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
