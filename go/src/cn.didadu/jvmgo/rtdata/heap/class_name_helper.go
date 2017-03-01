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