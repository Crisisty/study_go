package main

import (
	"testing"
)

func TestCh08(t *testing.T) {
	cmd := parseCmd()
	cmd.XjreOption = "/Library/Java/JavaVirtualMachines/jdk1.8.0_181.jdk/Contents/Home/jre"
	//cmd.class = "BubbleSortTest"
	//cmd.class = "HelloWorld"
	cmd.class = "PrintArgs"
	cmd.args = []string{"foo", "bar", "你好，世界！"}
	startJVM(cmd)
}
