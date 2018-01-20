package main

import (
	"log"
	"os"

	_ "code/chapter2/sample/matchers"
	"code/chapter2/sample/search"

	/*
	go编译器不允许声明导入却不使用。
	
	下划线让编译器接受这类导入（声明导入却不使用），并且调用对应包内的所有代码文件里定义的init函数。
	目的：调用 matchers 包中 rss.go 代码文件里的 init 函数，注册 RSS 匹配器，以便后用。
	*/
)

// init is called prior to main.
// init 在 main 之前调用
var i=0
func init() {
	// Change the device for logging to stdout.
	// 将日志输出到标准输出
	println("日志输出 + %d", i)
	log.SetOutput(os.Stdout)
	i += 1
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	//使用特定的项做搜索
	search.Run("president")
}
