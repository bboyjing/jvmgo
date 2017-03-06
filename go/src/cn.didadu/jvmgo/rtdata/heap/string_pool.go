package heap

import "unicode/utf16"

// 用map来表示字符串池，key是Go字符串，value是Java字符串
var internedStrings = map[string]*Object{}

// 根据Go自复查un返回相应的Java字符串实例
func JString(loader *ClassLoader, goStr string) *Object {
	// 如果在字符串池中存在，直接返回
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	// 将utf8转成utf16
	chars := stringToUtf16(goStr)
	// 字符串实例引用
	jChars := &Object{loader.LoadClass("[C"), chars, nil}

	// 加载String类，并且创建实例
	jStr := loader.LoadClass("java/lang/String").NewObject()
	// 通过反射，给实例的char[]类型的value变量设值
	jStr.SetRefVar("value", "[C", jChars)

	// 将字符串添加到常量池
	internedStrings[goStr] = jStr
	return jStr
}

// utf8 -> utf16
func stringToUtf16(s string) []uint16 {
	runes := []rune(s)         // utf32
	return utf16.Encode(runes) // func Encode(s []rune) []uint16
}


// java.lang.String -> go string
func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C")
	return utf16ToString(charArr.Chars())
}

// utf16 -> utf8
func utf16ToString(s []uint16) string {
	runes := utf16.Decode(s) // func Decode(s []uint16) []rune
	return string(runes)
}

func InternString(jStr *Object) *Object {
	goStr := GoString(jStr)
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	internedStrings[goStr] = jStr
	return jStr
}