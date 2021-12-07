package view

import (
	"fmt"
)

func Yellow(value interface{}) string {
	return fmt.Sprintf("\033[1;36m%v\033[0m", value)
}

func Teal(value interface{}) string {
	return fmt.Sprintf(" \033[1;33m%v\033[0m", value)
}