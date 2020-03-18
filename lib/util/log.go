package util

import (
	"fmt"
	"os"
)

// Fatalf prints the given format and args in red
// and then exits the process.
func Fatalf(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args...))
	os.Exit(1)
}

func Printf(format string, args ...interface{}) {
	fmt.Printf(fmt.Sprintf(format, args...))
}

func Printlnf(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args...))
}
