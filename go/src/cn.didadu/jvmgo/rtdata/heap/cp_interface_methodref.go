package heap

import "cn.didadu/jvmgo/classfile"

// 接口方法符号引用结构体
type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

// 初始化接口方法符号引用
func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

// 解析接口方法符号引用
func (self *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if self.method == nil {
		self.resolveInterfaceMethodRef()
	}
	return self.method
}

func (self *InterfaceMethodRef) resolveInterfaceMethodRef() {
	// 获取当前Class指针
	d := self.cp.class
	// 解析方法符号引用之前需要先解析方法所属的类
	c := self.ResolvedClass()
	// 判断方法所属类是否是接口
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 根据方法名和描述符查找方法
	method := lookupInterfaceMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	// 判断方法方法权限
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.method = method
}

func lookupInterfaceMethod(iface *Class, name, descriptor string) *Method {
	// 先从当前接口中查找方法
	for _, method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}

	// 若没有找到，则递归当前接口实现的接口链
	return lookupMethodInInterfaces(iface.interfaces, name, descriptor)
}