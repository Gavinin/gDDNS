package logger

var Log *Logger

type Logger interface {
	Debugf(format string, args any)
	Infof(format string, args any)
	Warnf(format string, args any)
	Warningf(format string, args any)
	Errorf(format string, args any)
	Fatalf(format string, args any)
	Panicf(format string, args any)
	Debug(args any)
	Print(args any)
	Warn(args any)
	Warning(args any)
	Error(args any)
	Fatal(args any)
	Panic(args any)
	Debugln(args any)
	Infoln(args any)
	Println(args any)
	Warnln(args any)
	Warningln(args any)
	Errorln(args any)
	Fatalln(args any)
	Panicln(args any)
	Info(args any)
	Printf(format string, args any)
}

func InitLog() {

}
