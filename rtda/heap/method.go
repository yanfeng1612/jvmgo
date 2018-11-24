package heap

type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	code      []byte
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].coypAttributes(cfMethod)
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
