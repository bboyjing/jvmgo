package heap

import (
	"cn.didadu/jvmgo/classfile"
	"strings"
)

type Class struct {
	accessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
}

// 将ClassFile结构体转换成Class结构体
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

// 判断访问标志符
func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

// self class是否能被other class访问
func (self *Class) isAccessibleTo(other *Class) bool {
	/*
		若self class访问标识为public或者两个类在同一个包下，则可以访问
		暂时简单地按照包名来检查类是否属于同一个包
	 */
	return self.IsPublic() ||
		self.getPackageName() == other.getPackageName()
}

// 获取包名
func (self *Class) getPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

// 判断self class是否时other class的子类
func (self *Class) isSubClassOf(other *Class) bool {
	return self.superClassName == other.name
}

// Getter方法
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}

// 创建对象引用
func (self *Class) NewObject() *Object {
	// 调用Object结构体
	return newObject(self)
}