package gosnow

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	logger   *log.Logger
	disabled = false
)

func init() {
	logger = log.New(os.Stdout, "logger: ", log.Lshortfile)
}

func DisableLogger() {
	disabled = true
	logger.SetFlags(0)
	SetLogOutput(ioutil.Discard)
}

func SetLogOutput(writer io.Writer) {
	logger.SetOutput(writer)
}

func GetWriter() io.Writer {
	return logger.Writer()
}
