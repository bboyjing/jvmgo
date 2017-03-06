package heap

import "cn.didadu/jvmgo/classfile"

type ExceptionTable []*ExceptionHandler

// 其中4个字段对应字节码中异常处理表的四个字段
type ExceptionHandler struct {
	// try{}语句块的第一条指令
	startPc   int
	// try{}语句块的下一条指令
	endPc     int
	handlerPc int
	// 若catchType为nil，则表示处理所有异常
	catchType *ClassRef
}

// 创建异常处理表
func newExceptionTable(entries []*classfile.ExceptionTableEntry, cp *ConstantPool) ExceptionTable {
	table := make([]*ExceptionHandler, len(entries))
	for i, entry := range entries {
		table[i] = &ExceptionHandler{
			startPc:   int(entry.StartPc()),
			endPc:     int(entry.EndPc()),
			handlerPc: int(entry.HandlerPc()),
			catchType: getCatchType(uint(entry.CatchType()), cp),
		}
	}

	return table
}

func getCatchType(index uint, cp *ConstantPool) *ClassRef {
	// index为0表示捕获所有异常
	if index == 0 {
		return nil // catch all
	}
	return cp.GetConstant(index).(*ClassRef)
}

// 查看异常处理表是否能处理异常
func (self ExceptionTable) findExceptionHandler(exClass *Class, pc int) *ExceptionHandler {
	for _, handler := range self {
		/*
			处理try{}语句块中抛出的异常
			try{]语句块包含startPc，但不包含endPc
		 */
		if pc >= handler.startPc && pc < handler.endPc {
			if handler.catchType == nil {
				return handler
			}
			catchClass := handler.catchType.ResolvedClass()
			if catchClass == exClass || catchClass.IsSuperClassOf(exClass) {
				return handler
			}
		}
	}
	return nil
}