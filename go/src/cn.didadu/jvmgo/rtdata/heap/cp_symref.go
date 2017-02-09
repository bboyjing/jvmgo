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