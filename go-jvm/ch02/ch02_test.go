package main

import (
	"testing"
)

func TestCh02(t *testing.T) {
	cmd := parseCmd()
	cmd.class = "Test"
	startJVM(cmd)
}
