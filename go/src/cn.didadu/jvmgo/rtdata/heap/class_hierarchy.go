package heap



// 判断other Class是不是self Class的实例
func (self *Class) isAssignableFrom(other *Class) bool {
	s, t := other, self

	// other Class和self Class是同一类型
	if s == t {
		return true
	}

	if !t.IsInterface() {
		// other Class是self Class的子类
		return s.isSubClassOf(t)
	} else {
		// other Class实现了self Class接口
		return s.isImplements(t)
	}
}

// 判断other class是否是self class的直接或间接超类
func (self *Class) isSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (self *Class) isImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

func (self *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}