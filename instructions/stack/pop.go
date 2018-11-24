package stack

import (
	"jvmgo/instructions/base"
	"jvmgo/rtda"
)

type POP struct{ base.NoOperandsInstruction }

type POP2 struct{ base.NoOperandsInstruction }

func (self *POP) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
}

func (self *POP2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
	frame.OperandStack().PopSlot()
}
