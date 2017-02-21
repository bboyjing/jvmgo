package heap

func LookupMethodInClass(class *Class, name, descriptor string) *Method {
	// 通过名称和描述符，从class以及superClass链中寻找方法
	for c := class; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}

func lookupMethodInInterfaces(ifaces []*Class, name, descriptor string) *Method {
	for _, iface := range ifaces {
		// 先遍历当前接口的方法
		for _, method := range iface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}

		// 如果没找到，再递归遍历当前接口实现的接口
		method := lookupMethodInInterfaces(iface.interfaces, name, descriptor)
		if method != nil {
			return method
		}
	}

	return nil
}