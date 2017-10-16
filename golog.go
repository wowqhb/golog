package golog

/**
ALL Level是最低等级的，用于打开所有日志记录。

DEBUG Level指出细粒度信息事件对调试应用程序是非常有帮助的。

INFO level表明 消息在粗粒度级别上突出强调应用程序的运行过程。

WARN level表明会出现潜在错误的情形。

ERROR level指出虽然发生错误事件，但仍然不影响系统的继续运行。

FATAL level指出每个严重的错误事件将会导致应用程序的退出。

OFF Level是最高等级的，用于关闭所有日志记录。
*/

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	ALL = 1 << iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	OFF
)

type GoLog struct {
	logger *log.Logger
	level  int
}

func NewGoLog(out io.Writer, level int) *GoLog {
	return &GoLog{
		logger: log.New(out, fmt.Sprintf("[PID: %d] ", os.Getpid()), log.LstdFlags),
		level:  level,
	}
}

func (this *GoLog) All(a ...interface{}) {
	if ALL >= this.level {
		this.logger.Println("ALL ", a)
	}
}

func (this *GoLog) Debug(a ...interface{}) {
	if DEBUG >= this.level {
		this.logger.Println("DEBUG ", a)
	}
}

func (this *GoLog) Info(a ...interface{}) {
	if INFO >= this.level {
		this.logger.Println("INFO ", a)
	}
}

func (this *GoLog) Warn(a ...interface{}) {
	if WARN >= this.level {
		this.logger.Println("WARN ", a)
	}
}

func (this *GoLog) Error(a ...interface{}) {
	if ERROR >= this.level {
		this.logger.Println("ERROR ", a)
	}
}

func (this *GoLog) Fatal(a ...interface{}) {
	if FATAL >= this.level {
		this.logger.Fatalln("FATAL ", a)
	}
}
