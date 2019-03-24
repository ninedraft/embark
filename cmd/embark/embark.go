package main

import (
	"github.com/ninedraft/embark/pkg/embark"
)

func main() {
	var err = embark.Main().Execute()
	if err != nil {
		embark.ErrPrintln(err)
	}
}
