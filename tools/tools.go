package tools

import (
	"fmt"
	"time"
)

func FechaMysql() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Second())
}
