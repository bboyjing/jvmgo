package heap

// 类、字段、方法和接口方法的符号引用公共结构体
type SymRef struct {
	// 运行时常量池指针，便于通过符号引用访问运行时常量池
	cp        *ConstantPool
	// 类的完全限定名
	className string
	// 缓存解析后的类结构体指针
	class     *Class
}

// 解析类符号引用
func (self *SymRef) ResolvedClass() *Class {
	// 如果类符号引用已经解析，直接返回类指针
	if self.class == nil {
		// 解析类符号
		self.resolveClassRef()
	}
	return self.class
}

func (self *SymRef) resolveClassRef() {
	// 获取当前Class指针
	d := self.cp.class
	// 通过需要引用的类的完全限定名加载类
	c := d.loader.LoadClass(self.className)
	// 判断d是否能有权限调用引用类c
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}