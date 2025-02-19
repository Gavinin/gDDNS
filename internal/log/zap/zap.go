package zap

import (
	zapl "go.uber.org/zap"
	"reflect"
	"time"
)

type Log struct {
	zap *zapl.SugaredLogger
}

func New() *Log {
	logger, err := zapl.NewProduction()
	if err != nil {
		panic(err)
	}

	return &Log{
		zap: logger.Sugar(),
	}
}

func (l *Log) Debugf(format string, args ...any) {
	l.zap.Debugf(format, args)
}

func (l *Log) Infof(format string, args ...any) {
	l.zap.Infof(format, args)
}

func (l *Log) Warnf(format string, args ...any) {
	l.zap.Warnf(format, args)
}

func (l *Log) Warningf(format string, args ...any) {
	l.zap.Warnf(format, args)
}

func (l *Log) Errorf(format string, args ...any) {
	l.zap.Errorf(format, args)
}

func (l *Log) Fatalf(format string, args ...any) {
	l.zap.Fatalf(format, args)
}

func (l *Log) Panicf(format string, args ...any) {
	l.zap.Panicf(format, args)
}

func (l *Log) Debug(args ...any) {
	l.zap.Debug(args)
}

func (l *Log) Print(args ...any) {
	l.zap.Info(args)
}

func (l *Log) Warn(args ...any) {
	l.zap.Warn(args)
}

func (l *Log) Warning(args ...any) {
	l.zap.Warn(args)
}

func (l *Log) Error(args ...any) {
	l.zap.Error(args)
}

func (l *Log) Fatal(args ...any) {
	l.zap.Fatal(args)
}

func (l *Log) Panic(args ...any) {
	l.zap.Panic(args)
}

func (l *Log) Debugln(args ...any) {
	l.zap.Debug(args)
}

func (l *Log) Infoln(args ...any) {
	l.zap.Info(args)
}

func (l *Log) Println(args ...any) {
	l.zap.Info(args)
}

func (l *Log) Warnln(args ...any) {
	l.zap.Warn(args)
}

func (l *Log) Warningln(args ...any) {
	l.zap.Warn(args)
}

func (l *Log) Errorln(args ...any) {
	l.zap.Error(args)
}

func (l *Log) Fatalln(args ...any) {
	l.zap.Fatal(args)
}

func (l *Log) Panicln(args ...any) {
	l.zap.Panic(args)
}

func (l *Log) Info(args ...any) {
	l.zap.Info(args)
}

func (l *Log) Printf(format string, args ...any) {
	l.zap.Infof(format, args)
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
