package heap

import (
	"jvmgo/classfile"
)

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(refINfo)
	return ref
}

func (self *InterfaceMethodRef) ResolveInterfaceMethod() *Method {
	if self.method == nil {
		self.ResolveInterfaceMethodRef()
	}
	return self.method
}

func (self *InterfaceMethodRef) ResolveInterfaceMethodRef() *Method {
	d := self.cp.class
	c := self.ResolveClass()
	if !c.IsInterface() {
		panic("java.lang.IncopatibleClassChangeError")
	}
	method := lookupInterfaceMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.method = method
}

func lookupInterfaceMethod(iface *Class, name string, descriptor string) {
	for _, method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return lookupMethodInInterfaces(iface.interfaces, name, descriptor)
}
