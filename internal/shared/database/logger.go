package database

import (
	"fmt"
	"os"
)

type stdoutWriter struct{}

func (w stdoutWriter) Printf(format string, args ...interface{}) {
	_, err := fmt.Fprintf(os.Stdout, format, args...)
	if err != nil {
		return
	}
}
