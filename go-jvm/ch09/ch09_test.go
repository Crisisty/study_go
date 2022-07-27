package main

import (
	"testing"
)

func TestCh09(t *testing.T) {
	cmd := parseCmd()
	cmd.XjreOption = "/Library/Java/JavaVirtualMachines/jdk1.8.0_181.jdk/Contents/Home/jre"
	//cmd.class = "GetClassTest"
	//cmd.class = "StringTest"
	//cmd.class = "ObjectTest"
	//cmd.class = "CloneTest"
	cmd.class = "BoxTest"
	startJVM(cmd)
}
