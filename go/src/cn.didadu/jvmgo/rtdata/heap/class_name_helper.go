package heap

/*
	将类名转换成类型描述符
	XXX -> [LXXX;
	int -> [I
	[XXX -> [[XXX
 */
func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}

/*
	将类名转成类型描述符
	XXX  => LXXX;
	int  => I
	[XXX => [XXX
 */
func toDescriptor(className string) string {
	// 如果已经是数组类名，直接返回
	if className[0] == '[' {
		// array
		return className
	}
	// 如果是基本类型名，返回对应的类型描述符
	if d, ok := primitiveTypes[className]; ok {
		// primitive
		return d
	}
	// 普通类名转成类型描述符L***;
	return "L" + className + ";"
}

// 基本类型和对应类型描述符map
var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

/*
	将数组类名转换成类名
	[[XXX -> [XXX
	[LXXX; -> XXX
	[I -> int
 */
func getComponentClassName(className string) string {
	if className[0] == '[' {
		componentTypeDescriptor := className[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array: " + className)
}

// [XXX  => [XXX
// LXXX; => XXX
// I     => int
func toClassName(descriptor string) string {
	// 若是数组类，直接返回
	if descriptor[0] == '[' {
		// array
		return descriptor
	}
	// 若是引用类型，去掉前缀L
	if descriptor[0] == 'L' {
		// object
		return descriptor[1 : len(descriptor) - 1]
	}
	// 若是基本类型，通过基本类型描述符获取基本类型
	for className, d := range primitiveTypes {
		if d == descriptor {
			// primitive
			return className
		}
	}
	panic("Invalid descriptor: " + descriptor)
}