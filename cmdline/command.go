package cmdline

import (
	"fmt"
)

// java [options] class [arg arg arg]
type Command struct {
	options *Options
	class   string
	args    []string
}

func (self *Command) Options() *Options {
	return self.options
}

func (self *Command) Class() string {
	return self.class
}

func (self *Command) Args() []string {
	return self.args
}

func ParseCommand(osArgs []string) (cmd *Command, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	argReader := &ArgReader{osArgs[1:]}
	cmd = &Command{
		options: parseOptions(argReader),
		class:   argReader.removeFirst(),
		args:    argReader.args,
	}
	return
}

func PrintUseage() {
	fmt.Println("useage: jvmgo [-options] class [args ...]")
}
