package heap

import (
	"jvmgo/classfile"
)

type Method struct {
	ClassMember
	maxStack      uint
	maxLocals     uint
	code          []byte
	argsSlotCount uint
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].coypAttributes(cfMethod)
		methods[i].argsSlotCount = calcArgsSlotCount(methods[i])
	}
	return methods
}

func (self *Method) coypAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttrbute(); codeAttr != nil {
		self.maxStack = codeAttr.maxLocals
		self.maxLocals = codeAttr.maxLocals
		self.code = codeAttr.code
	}
}

func (self *Method) ArgsSlotCount() uint {
	return self.argsSlotCount
}

func calcArgsSlotCount(method *Method) uint {
	var result uint
	methodDescriptor := parseMethodDescriptor(method.descriptor)
	for _, paramType := range methodDescriptor.parameterTypes {
		result++
		if paramType == "J" || paramType == "D" {
			result++
		}
	}
	if !method.IsStatic {
		result++
	}
	return result
}

func parseMethodDescriptor(descriptor string) MethodDescriptor {

}
