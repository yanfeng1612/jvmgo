package base

type BytecodeReader struct {
	code []byte
	pc   int
}

func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

func (self *BytecodeReader) ReadUint8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}

func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

func (self *BytecodeReader) ReadUint16() uint16 {
	high := uint16(self.code[self.pc])
	low := uint16(self.code[self.pc+1])
	self.pc += 2
	return (high << 8) | low
}

func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

func (self *BytecodeReader) ReadInt32() int32 {
	high := int32(self.ReadInt16())
	low := int32(self.ReadInt16())
	return high<<16 | low
}
