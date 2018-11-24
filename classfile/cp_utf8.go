package classfile

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	n := uint32(reader.readUint16())
	bytes := reader.readBytes(n)
	self.str = string(bytes)
}
