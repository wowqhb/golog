package main

import (
	"github.com/wowqhb/golog"
	"os"
)

func main() {
	logger := golog.NewGoLog(os.Stdout, golog.DEBUG)
	logger.Info("sssss")

	file, _ := os.Create("./log.log")

	loggerFile := golog.NewGoLog(file, golog.INFO)
	loggerFile.Info("sssss")
}
