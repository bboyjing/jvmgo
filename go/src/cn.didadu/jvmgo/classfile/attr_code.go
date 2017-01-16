package classfile

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

// 读取Code属性
func (self *CodeAttribute) readInfo(reader *ClassReader) {
	// 读取2个字节的maxStack
	self.maxStack = reader.readUint16()
	// 读取2个字节的maxLocals
	self.maxLocals = reader.readUint16()
	// 读取4个字节的code长度
	codeLength := reader.readUint32()
	// 读取指定字节数的code
	self.code = reader.readBytes(codeLength)
	// 读取异常处理表
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

// 读取异常处理表
func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	// 读取2个字节的exception_table_length
	exceptionTableLength := reader.readUint16()
	// 读取异常处理表
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}

// 将maxStack转换成int
func (self *CodeAttribute) MaxStack() uint {
	return uint(self.maxStack)
}

// 将maxLocals转换成int
func (self *CodeAttribute) MaxLocals() uint {
	return uint(self.maxLocals)
}

func (self *CodeAttribute) Code() []byte {
	return self.code
}
func (self *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return self.exceptionTable
}

/*
	读取异常处理表相关信息
 */
func (self *ExceptionTableEntry) StartPc() uint16 {
	return self.startPc
}
func (self *ExceptionTableEntry) EndPc() uint16 {
	return self.endPc
}
func (self *ExceptionTableEntry) HandlerPc() uint16 {
	return self.handlerPc
}
func (self *ExceptionTableEntry) CatchType() uint16 {
	return self.catchType
}