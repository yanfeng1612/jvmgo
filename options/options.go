package options

import (
	"jvmgo/cmdline"
	"os"
	"path/filepath"
)

var (
	VerboseClass    bool
	ThreadStackSize uint
	AbsJavaHome     string
	AbsJreLib       string
)

func InitOptions(cmdOptions *cmdline.Options) {
	VerboseClass = cmdOptions.VerboseClass()
	ThreadStackSize = uint(cmdOptions.Xss)
	initJavaHome(cmdOptions.XuseJavaHome)
}

func initJavaHome(useOsEnv bool) {
	jh := "./jre"
	if useOsEnv {
		jh = os.Getenv("JAVA_HOME")
		if jh == "" {
			panic("$JAVA_HOME not set!")
		}
	}
	if absJh, err := filepath.Abs(jh); err == nil {
		AbsJavaHome = absJh
		AbsJreLib = filepath.Join(absJh, "lib")
	} else {
		panic(err)
	}
}
