package cfmt

import (
	"fmt"
	"testing"
)

func TestCfmt(t *testing.T) {
	Info("unit ", "test: info")
	err := ErrorE("unit test ", "error")
	fmt.Println(err)
}
