package main

import (
	"testing"
)

func TestCh11(t *testing.T) {
	cmd := parseCmd()
	cmd.XjreOption = "/Library/Java/JavaVirtualMachines/jdk1.8.0_181.jdk/Contents/Home/jre"
	//cmd.args = []string{"abc"}
	cmd.class = "HelloWorld"
	newJVM(cmd).start()
}
