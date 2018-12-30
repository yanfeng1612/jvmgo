package heap

import (
	"jvmgo/classfile"
)

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MemberRef {
	ref := &MemberRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMethodrefInfo)
	return ref
}
func (self *MemberRef) ResolveMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

func (self *MethodRef) resovleMethodRef() {
	d := self.cp.class
	c := self.ResolveClass()
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessorError")
	}
	self.method = method
}

func lookupMethod(class *Class, name string, descriptor string) *Method {
	method := lookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
