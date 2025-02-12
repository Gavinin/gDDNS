package zap

import (
	zapl "go.uber.org/zap"
	"reflect"
	"time"
)

type Log struct {
	zap *zapl.Logger
}

func New() *Log {
	logger, err := zapl.NewProduction()
	if err != nil {
		panic(err)
	}
	return &Log{
		zap: logger,
	}
}

func (l *Log) Debugf(format string, args any) {
	l.zap.Debug(format, args)
}

func (l *Log) Infof(format string, args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Warnf(format string, args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Warningf(format string, args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Errorf(format string, args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Fatalf(format string, args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Panicf(format string, args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Debug(args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Print(args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Warn(args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Warning(args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Error(args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Fatal(args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Panic(args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Debugln(args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Infoln(args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Println(args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Warnln(args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Warningln(args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Errorln(args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Fatalln(args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Panicln(args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Info(args any) {
	//TODO implement me
	panic("implement me")
}

func (l *Log) Printf(format string, args any) {
	//TODO implement me
	panic("implement me")
}
func logKeyValue(key string, value any) zapl.Field {
	switch reflect.TypeOf(value).Kind() {
	case reflect.String:
		return zapl.String(key, value.(string))
	case reflect.Int, reflect.Int32, reflect.Int64:
		return zapl.Int64(key, value.(int64))
	case reflect.TypeOf(time.Duration(0)).Kind():
		return zapl.Duration(key, value.(time.Duration))
	default:
		return zapl.Any(key, value)
	}
}
