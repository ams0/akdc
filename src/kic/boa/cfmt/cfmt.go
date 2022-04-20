package cfmt

import (
	"fmt"
	"os"
	"runtime"
)

var (
	Reset  string
	Blue   string
	Cyan   string
	Gray   string
	Green  string
	Purple string
	Red    string
	White  string
	Yellow string
)

// print an info message in cyan
func Info(msg ...interface{}) {
	fmt.Print(Cyan)
	fmt.Print(msg...)
	fmt.Println(Reset)
}

// print the error in red and return params as error
func ErrorE(msg ...interface{}) error {
	fmt.Print(Red)
	fmt.Print(msg...)
	fmt.Println(Reset)
	return fmt.Errorf("%v", fmt.Sprint(msg...))
}

func init() {
	// Windows doesn't support ANSI colors
	if runtime.GOOS != "windows" {
		Reset = "\033[0m"
		Blue = "\033[34m"
		Cyan = "\033[36m"
		Gray = "\033[37m"
		Green = "\033[32m"
		Purple = "\033[35m"
		Red = "\033[31m"
		White = "\033[97m"
		Yellow = "\033[33m"
	}
}

// append text to path
func FAppendToFile(path string, text string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return ErrorE(err)
	}

	defer f.Close()

	if _, err := f.WriteString(text); err != nil {
		return ErrorE(err)
	}

	return nil
}
