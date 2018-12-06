package references

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type NEW struct {
	base.Index16Instruction
}

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
}
