package embark

import (
	"fmt"
	"os"
)

func ErrPrint(args ...interface{}) {
	fmt.Fprint(os.Stderr, args...)
}

func ErrPrintf(ff string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, ff, args...)
}

func ErrPrintln(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}
