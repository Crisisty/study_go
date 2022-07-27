package main

import (
	"fmt"
	"jvmgo/ch03/classfile"
)
import "strings"
import "jvmgo/ch03/classpath"

func main() {
	terminal := parseCmd()
	if terminal.versionFlag {
		fmt.Println("version 0.0.1")
	} else if terminal.helpFlag || terminal.class == "" {
		printUsage()
	} else {
		startJVM(terminal)
	}
}

func startJVM(terminal *Cmd) {
	cp := classpath.Parse(terminal.XjreOption, terminal.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n",
		cp, terminal.class, terminal.args)
	className := strings.Replace(terminal.class, ".", "/", -1)
	cf := loadClass(className, cp)
	printClassInfo(cf)
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access fields: %v\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SupperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf(" %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))

	for _, m := range cf.Methods() {
		fmt.Printf(" %s\n", m.Name())
	}
}
