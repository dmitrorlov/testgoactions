package main

import (
	"github.com/dmitrorlov/testgoactions/src"
	"github.com/goutlz/logz"
)

func main() {
	l := logz.NewColoredLogger()
	l.Infof("START")
	l.Infof("SUM: %v", src.Sum(3, 2))
	l.Infof("MULTIPLY: %v", src.Multiply(3, 2))
	l.Infof("END")
}
