package heap

import (
	"cn.didadu/jvmgo/classpath"
	"fmt"
	"cn.didadu/jvmgo/classfile"
)

type ClassLoader struct {
	// 通过Classpath来搜索和读取class文件
	cp          *classpath.Classpath
	// 记录已经加载的类数据，key是类完全限定名
	classMap    map[string]*Class
	// 是否把类加载信息输出到控制台
	verboseFlag bool
}

// 创建ClassLoader实例
func NewClassLoader(cp *classpath.Classpath, verboseFlag bool) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		verboseFlag: verboseFlag,
		classMap: make(map[string]*Class),
	}
}

// 把类数据加载到方法区
func (self *ClassLoader) LoadClass(name string) *Class {
	// 判断类是否已经被加载
	if class, ok := self.classMap[name]; ok {
		// 已加载直接返回
		return class
	}

	// 若类名的第一个字符是'['，表示该类是数组类
	if name[0] == '[' {
		return self.loadArrayClass(name)
	}

	// 加载类
	return self.loadNonArrayClass(name)
}

// 加载数组类
func (self *ClassLoader) loadArrayClass(name string) *Class {
	/*
		生成Class结构体，数组类由Java虚拟机在运行时生成
		所以没有像loadNonArrayClass()那样读取class文件
	 */
	class := &Class{
		// 访问标识
		accessFlags: ACC_PUBLIC, // todo
		// 类名
		name:        name,
		// 类加载器
		loader:      self,
		// 数组类不需要初始化，所以设成true
		initStarted: true,
		// 数组类的超类是java/lang/Object
		superClass:  self.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			// 数组类实现了java.lang.Cloneable接口
			self.LoadClass("java/lang/Cloneable"),
			// 数组类实现了java.io.Serializable
			self.LoadClass("java/io/Serializable"),
		},
	}
	// 记录该数组类已加载
	self.classMap[name] = class
	return class
}

/*
	数组类和普通类有很大的不同
	它的数据并不是来自class文件，而是由Java虚拟机在运行期间生成
	暂不考虑数组类的加载
 */
func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	// 找到class文件并把数据读到内存
	data, entry := self.readClass(name)
	// 解析class文件，生成虚拟机可以使用的类数据
	class := self.defineClass(data)
	// 链接
	link(class)

	if self.verboseFlag {
		fmt.Printf("[Loaded %s from %s]\n", name, entry)
	}

	return class
}

// 读取class
func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	// 调用Classpath的ReadClass()方法
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		// 若出错则跑出ClassNotFoundException异常
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

// 解析class
func (self *ClassLoader) defineClass(data []byte) *Class {
	// 从byte数组中获取Class结构体
	class := parseClass(data)
	// 设置Class结构它的classloader指针
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

// 把class文件数据转换成Class结构体
func parseClass(data []byte) *Class {
	// 从byte数组中读取ClassFile
	cf, err := classfile.Parse(data)
	if err != nil {
		//panic("java.lang.ClassFormatError")
		panic(err)
	}
	// 返回class结构体
	return newClass(cf)
}

// 解析超类符号引用
func resolveSuperClass(class *Class) {
	// 除了java.lang.Object，否则要递归调用LoadClass()方法加载超类
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

// 解析接口符号引用
func resolveInterfaces(class *Class) {
	// 获取接口个数
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

// 链接
func link(class *Class) {
	verify(class)
	prepare(class)
}

// 验证阶段
func verify(class *Class) {
	// todo
}

// 准备阶段
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

// 计算实例字段的个数，同时编号
func calcInstanceFieldSlotIds(class *Class) {
	// 初始化slotId为0
	slotId := uint(0)
	// 如果超类不为空，slotId从超类的字段个数后开始编号
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		// 判断是否是实例字段
		if !field.IsStatic() {
			// 设置字段对应Slots数组中的位置
			field.slotId = slotId
			slotId++
			// 若字段类型是long或double，则占2个slot
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	// 设置Class的实例字段个数
	class.instanceSlotCount = slotId
}

// 计算静态字段的个数，同时编号
func calcStaticFieldSlotIds(class *Class) {
	// 初始化slotId为0
	slotId := uint(0)
	for _, field := range class.fields {
		// 判断是否是静态变量
		if field.IsStatic() {
			// 设置字段对应Slots数组中的位置
			field.slotId = slotId
			slotId++
			// 若字段类型是long或double，则占2个slot
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	// 设置Class的静态字段个数
	class.staticSlotCount = slotId
}

// 给静态变量分配空间，同时赋予初始值
func allocAndInitStaticVars(class *Class) {
	/*
		初始化Class静态变量
		Go语言保证新创建的Slot结构体有默认值：num字段为0，ref字段为nil
	 */
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		// 判断是否是私有的静态变量
		if field.IsStatic() && field.IsFinal() {
			// 初始化静态私有变量
			initStaticFinalVar(class, field)
		}
	}
}

// 初始化类静态私有变量
func initStaticFinalVar(class *Class, field *Field) {
	// 获取Class的静态数组
	vars := class.staticVars
	// 获取运行时常量池
	cp := class.constantPool
	// 获取常量值索引(在常量池中第几个)
	cpIndex := field.ConstValueIndex()
	// 获取字段在静态数组中的位置
	slotId := field.SlotId()

	if cpIndex > 0 {
		/*
			根据字段描述符(字段类型)读取相应的值
			并将读取的常量值赋到Class静态数组对应的位置
		 */
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		// String类型
		case "Ljava/lang/String;":
			// 获取常量池中存储的Go自复查un
			goStr := cp.GetConstant(cpIndex).(string)
			// 获取字符串池中存储的Java字符串
			jStr := JString(class.Loader(), goStr)
			vars.SetRef(slotId, jStr)
		}
	}
}