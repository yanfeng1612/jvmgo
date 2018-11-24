package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operantStack *OperandStack
}

func newFrame(maxLocals uint, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operantStack: newOperandStack(maxStack),
	}
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operantStack
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
