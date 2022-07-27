package main

import (
	"jvmgo/ch05/classpath"
	"strings"
	"testing"
)

func TestCh05(t *testing.T) {
	cp := classpath.Parse("", "")
	className := strings.Replace("GaussTest", ".", "/", -1)
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		//fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}
