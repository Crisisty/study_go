package classpath

import (
	"fmt"
	"os"
	"path/filepath"
)

type Classpath struct {
	//启动类加载器
	bootClasspath Entry
	//扩展类加载器
	extClasspath Entry
	//应用程序加载器
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	fmt.Printf("%s\n", jreOption)
	fmt.Printf("%s\n", cpOption)
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	fmt.Printf("className:%s\n", className)
	if self.bootClasspath != nil {
		if data, entry, err := self.bootClasspath.readClass(className); err == nil {
			fmt.Println("bootClasspath.readClass")
			fmt.Println(entry.String())
			return data, entry, err
		}
	}
	if self.extClasspath != nil {
		if data, entry, err := self.extClasspath.readClass(className); err == nil {
			fmt.Println("extClasspath.readClass")
			return data, entry, err
		}
	}
	fmt.Println("userClasspath.readClass")
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	if jreOption == "" {
		return
	}
	fmt.Println("启动类加载器启动")
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	fmt.Printf("boot=%s\n", self.bootClasspath.String())
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
	fmt.Printf("ext=%s\n", self.extClasspath.String())
}

func (self *Classpath) parseUserClasspath(cpOption string) {
	fmt.Println("应用程序类加载器启动")
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
	fmt.Printf("user=%s\n", self.userClasspath.String())
}

func getJreDir(jreOption string) string {
	//优先使用jreOption
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func exists(path string) bool {
	//检出路径的状态，是否存在异常
	if _, err := os.Stat(path); err != nil {
		//存在异常
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
