package cfmt

import (
	"fmt"
	"os"
	"os/exec"
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
	fmt.Println(fmt.Sprintf("%v", msg))
	fmt.Print(Reset)
}

func Error(msg ...interface{}) {
	fmt.Print(Red)
	fmt.Println(msg)
	fmt.Print(Reset)
}

func ExitError(err error) {
	exitError := err.(*exec.ExitError)
	Error(exitError)
	os.Exit(2)
}

func init() {
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
