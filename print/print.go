package print

import (
	"fmt"
	"os"
)

var VerboseEnabled bool

func init() {
	VerboseEnabled = false
}

func Verbose(msg string) {
	if VerboseEnabled {
		fmt.Print(msg)
	}
}

func Verboseln(msg string) {
	if VerboseEnabled {
		fmt.Println(msg)
	}
}

func Verbosef(msg string, args ...interface{}) {
	if VerboseEnabled {
		fmt.Printf(msg, args...)
	}
}

func Info(msg string) {
	fmt.Print(msg)
}

func Infoln(msg string) {
	fmt.Println(msg)
}

func Infof(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}

func Fatal(msg string) {
	fmt.Print(msg)
	os.Exit(1)
}

func Fatalln(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func Fatalf(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
	os.Exit(1)
}
