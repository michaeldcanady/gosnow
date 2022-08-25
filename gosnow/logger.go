package gosnow

import (
	"io"
	"log"
	"os"
)

var (
	logger *log.Logger
	//disabled = false
)

func init() {
	logger = log.New(os.Stdout, "gosnow - ", log.Lshortfile)
}

// DisableLogger disables the built-in module logger
func DisableLogger() {
	//disabled = true
	logger.SetFlags(0)
	SetLogOutput(io.Discard)
}

// SetLogOutput used to change where the logger writes to
func SetLogOutput(writer io.Writer) {
	logger.SetOutput(writer)
}

// GetWriter returns the writer
func GetWriter() io.Writer {
	return logger.Writer()
}
