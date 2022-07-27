package main

import (
	"fmt"
	"jvmgo/ch07/classpath"
	"jvmgo/ch07/rtda/heap"
	"strings"
	"testing"
)

func TestCh07(t *testing.T) {
	cp := classpath.Parse("/Library/Java/JavaVirtualMachines/jdk1.8.0_181.jdk/Contents/Home/jre", "")
	classloader := heap.NewClassLoader(cp, true)
	className := strings.Replace("FibonacciTest", ".", "/", -1)
	mainClass := classloader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod, true)
	} else {
		fmt.Printf("Main method not found in class %s\n", "MyObject")
	}
}
