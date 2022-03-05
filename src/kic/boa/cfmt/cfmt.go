package cfmt

import (
	"fmt"
	"os"
	// "os/exec"
	"runtime"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

// todo - adding [] to output
func Info(msg ...interface{}) {
	fmt.Print(Cyan)
	fmt.Printf("%v", msg)
	fmt.Println(Reset)
}

func Error(msg ...interface{}) {
	fmt.Print(Red)
	fmt.Printf("%v", msg)
	fmt.Println(Reset)
}

func ExitError(err error) {
	//msg := err.(*exec.ExitError)
	Error(err)
	os.Exit(1)
}

func ExitErrorMessage(msg ...interface{}) {
	Error(msg)
	os.Exit(1)
}

func init() {
	// Windows doesn't support ANSI colors
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
}
