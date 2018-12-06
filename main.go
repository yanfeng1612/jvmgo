package main

import (
	"jvmgo/cmdline"
	"jvmgo/options"
	"os"
	"runtime/pprof"
)

func main() {
	cmd, err := cmdline.ParseCommand(os.Args)
	if err != nil {
		cmdline.PrintUseage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *cmdline.Command) {
	Xcpuprofile := cmd.Options().Xcpuprofile
	if Xcpuprofile != "" {
		f, err := os.Create(Xcpuprofile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	options.InitOptions(cmd.Options())

	// cp := classpath.Parse(cmd.Options().Classpath())
	// heap.InitBootLoader(cp)

	// mainClassName := strings.Replace(cmd.Class(), ".", "/", -1)
	// mainThread := createMainThread(mainClassName, cmd.Args())
	// interpreter.Loop(mainThread)
	// interpreter.KeepAlive()
}

// func createMainThread(className string, args []string) *rtda.Thread {
// 	mainThread := rtda.NewThread(nil)
// 	bootMethod := heap.BootstrapMethod()
// 	bootArgs := []interface{}{className, args}
// 	mainThread.InvokeMethodWithShim(bootMethod, bootArgs)
// 	return mainThread
// }
