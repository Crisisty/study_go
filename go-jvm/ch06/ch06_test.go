package main

import (
	"fmt"
	"jvmgo/ch06/classpath"
	"jvmgo/ch06/rtda/heap"
	"strings"
	"testing"
)

func TestCh06(t *testing.T) {
	cp := classpath.Parse("/Library/Java/JavaVirtualMachines/jdk1.8.0_181.jdk/Contents/Home/jre", "")
	classloader := heap.NewClassLoader(cp)
	className := strings.Replace("MyObject", ".", "/", -1)
	mainClass := classloader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", "MyObject")
	}
}
