package cmdline

const (
	_1k = 1024
	_1m = _1k * _1k
	_1g = _1m * _1m
)

type Options struct {
	classpath    string
	verboseClass bool
	Xss          int
	Xcpuprofile  string
	XuseJavaHome string
}

func parseOptions(argReader *ArgReader) *Options {
	option := &Options{
		Xss: 64 * _1k,
	}
	for argReader.hasMoreOptions() {
		optionName := argReader.removeFirst()
		switch optionName {
		case "-cp", "-classpath":
			option.classpath = argReader.removeFirst()
		case "-verbose":
			option.verboseClass = true
		case "-Xcpuprofile":
			option.Xcpuprofile = argReader.removeFirst()
		case "-XuseJavaHome":
			option.XuseJavaHome = argReader.removeFirst()
		}
	}
	return option
}
