package classfile

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}
