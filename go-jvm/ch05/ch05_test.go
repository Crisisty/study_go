package main

import (
	"testing"
)

func TestCh05(t *testing.T) {
	cmd := parseCmd()
	cmd.class = "GaussTest"
	startJVM(cmd)
}
