package database

import (
	"fmt"
	"os"
)

type stdoutWriter struct{}

func (w stdoutWriter) Printf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, format, args...)
}
