package gosnow

import (
	"io"
	"log"
	"os"
)

var (
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stdout, "logger: ", log.Lshortfile)
}

func SetLogOutput(writer io.Writer) {
	logger.SetOutput(writer)
}

func GetWriter() io.Writer {
	return logger.Writer()
}
