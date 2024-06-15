package mylogger

import (
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"runtime"
	"time"
)

type Level logger.LogLevel

const (
	LevelDebug logger.LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

type Fields map[string]interface{}

func String(level logger.LogLevel) string {
	switch level {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	default:
		return "error"
	}
}

type LoggerImp struct {
	loglevel  Level
	newLogger *log.Logger
	ctx       context.Context
	fields    Fields
	callers   []string
}

func (l *LoggerImp) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.loglevel = Level(level)
	return &newlogger
}

func (l *LoggerImp) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	//TODO implement me
	panic("implement me")
}

func NewLogger(w io.Writer, prefix string, flag int) *LoggerImp {
	l := log.New(w, prefix, flag)
	return &LoggerImp{newLogger: l}
}

func (l *LoggerImp) clone() *LoggerImp {
	newlog := *l
	return &newlog
}

//WithFields 设置公共字段

func (l *LoggerImp) WithFields(f Fields) *LoggerImp {
	newlog := l.clone()
	if newlog.fields == nil {
		newlog.fields = make(Fields)
	}
	for k, v := range f {
		newlog.fields[k] = v
	}
	return newlog
}

//WithContext 设置日志上下文属性

func (l *LoggerImp) WithContext(ctx context.Context) *LoggerImp {
	newlog := l.clone()
	newlog.ctx = ctx
	return newlog
}

// WithCaller 设置当前某一层(skip) 调用栈的信息(对应函数信息，文件信息，行号)

func (l *LoggerImp) WithCaller(skip int) *LoggerImp {
	newlog := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		newlog.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}
	return newlog
}

// WithCallersFrames 设置记录整个调用栈信息

func (l *LoggerImp) WithCallersFrames() *LoggerImp {
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		info := fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function)
		callers = append(callers, info)
		if !more {
			break
		}
	}
	newlog := l.clone()
	newlog.callers = callers
	return newlog
}

// JSONFormat 将要输出的key-value整合到一个map中

func (l *LoggerImp) JSONFormat(level logger.LogLevel, message string) map[string]interface{} {
	data := make(Fields, len(l.fields)+4)
	data["level"] = String(level)
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			data[k] = v
		}
	}
	return data
}

//Output 将对应level级别内容进行log记录写入磁盘

func (l *LoggerImp) Output(ctx context.Context, level logger.LogLevel, message string) {
	marshal, _ := json.Marshal(l.JSONFormat(level, message))
	content := string(marshal)
	switch level {
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	}
}

// 各类级别信息级别的日志输出，外部调用

func (l *LoggerImp) Info(ctx context.Context, s string, v ...interface{}) {
	l.Output(ctx, LevelInfo, s+", "+fmt.Sprint(v...))
}

func (l *LoggerImp) Infof(ctx context.Context, format string, v ...interface{}) {
	l.Output(ctx, LevelInfo, fmt.Sprintf(format, v...))
}

func (l *LoggerImp) Debug(ctx context.Context, s string, v ...interface{}) {
	l.Output(ctx, LevelInfo, s+", "+fmt.Sprint(v...))
}

func (l *LoggerImp) Debugf(ctx context.Context, format string, v ...interface{}) {
	l.Output(ctx, LevelInfo, fmt.Sprintf(format, v...))
}

func (l *LoggerImp) Warn(ctx context.Context, s string, v ...interface{}) {
	l.Output(ctx, LevelInfo, s+", "+fmt.Sprint(v...))
}

func (l *LoggerImp) Warnf(ctx context.Context, format string, v ...interface{}) {
	l.Output(ctx, LevelInfo, fmt.Sprintf(format, v...))
}

func (l *LoggerImp) Error(ctx context.Context, s string, v ...interface{}) {
	l.Output(ctx, LevelInfo, s+", "+fmt.Sprint(v...))
}

func (l *LoggerImp) Errorf(ctx context.Context, format string, v ...interface{}) {
	l.Output(ctx, LevelInfo, fmt.Sprintf(format, v...))
}

func (l *LoggerImp) Fatal(ctx context.Context, v ...interface{}) {
	l.Output(ctx, LevelInfo, fmt.Sprint(v...))
}

func (l *LoggerImp) Fatalf(ctx context.Context, format string, v ...interface{}) {
	l.Output(ctx, LevelInfo, fmt.Sprintf(format, v...))
}

func (l *LoggerImp) Panic(ctx context.Context, v ...interface{}) {
	l.Output(ctx, LevelInfo, fmt.Sprint(v...))
}

func (l *LoggerImp) Panicf(ctx context.Context, format string, v ...interface{}) {
	l.Output(ctx, LevelInfo, fmt.Sprintf(format, v...))
}
