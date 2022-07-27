package main

import (
	"fmt"
	"testing"
)

func TestCh01(t *testing.T) {
	cmd := parseCmd()
	//cmd.versionFlag = true
	cmd.helpFlag = true
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}
