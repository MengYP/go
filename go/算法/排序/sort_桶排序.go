// 桶排序：最快最简单的排序
package main

import (
	"os"
	"fmt"
)

var (
	firstName, lastName string
	i , j int
	f float32
	input = "56.12 / 5252 / Go"
	format = "%f / %d / %s"
)
type point struct {
	x, y int
}

func main(){
	// fmtPrint()
	simpleBucketSort()
	


}



func simpleBucketSort() {
	
	var inputA int
	var sortA [11]int 
	var k int
	for k = 0; k < 5; k++ {
		fmt.Scanln(&inputA)
		if inputA < 11 {
			sortA[inputA] ++
		}else{
			fmt.Printf("输入数字必须是0~10\n")
			k--
		}
	}
	fmt.Printf("-----\n")
	fmt.Printf("桶排序：\n")
	for j = 0; j < len(sortA); j++ {
		sortValue := sortA[j]
		if sortValue != 0 {
			fmt.Printf("sortA[%d] = %d\n", j, sortA[j] )
			fmt.Printf("sortA[%d]: ",j)
			for index := 0; index < sortValue; index++ {
				fmt.Printf(" %d ",j)
			}
			fmt.Printf("\n----\n")
		}
	}
}

func fmtPrint() {
	fmt.Println("Please input your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi, %s %s !\n", firstName, lastName)

	// go格式化输出：
	// %v : 结构体的对象的值
	// %+v :结构体的成员名称和值 
	// %#v :go语法表示方式
	// %T : 输出一个值的数据类型
	p := point{1, 2}
	fmt.Printf("%v\n", p)  // {1 2}
	fmt.Printf("%+v\n", p) // {x:1 y:2}
	fmt.Printf("%#v\n", p) // main.point{x:1, y:2}
	fmt.Printf("%T\n",p)   // main.point
	// %t :格式化布尔型变量
	fmt.Printf("%t\n", true) // true
	// %d :以10进制来输出整型的方式
	fmt.Printf("%d\n", 123) // 123
	// %b :输出整型的二进制表示方式
	fmt.Printf("%b\n", 123) // 1111011
	// %c :打印出该整型数值所对应的字符
	fmt.Printf("%c\n", 33) // !
	// %x :输出一个值的16进制表示方式 
	fmt.Printf("%x\n", 456) // 1c8
	//%f : 格式化浮点型数值
	//%e %E :使用科学计数法来输出整型
	fmt.Printf("%f\n", 78.8) //78.800000
	fmt.Printf("%e\n", 12000000000.0) // 1.200000e+10
	fmt.Printf("%E\n", 12000000000.0) // 1.200000E+10
	//%s : 输出基本的字符串
	//%q : 输出像Go源码中那样带双引号的字符串，需使用`%q`
	//%x : 以16进制输出字符串，每个字符串的字节用两个字符输出
	fmt.Printf("%s\n","\"string\"") // "string"
	fmt.Printf("%q\n","\"string\"") // "\"string\""
	fmt.Printf("%x\n", "Meng") // 4d656e67
	//%p : 输出一个指针的值
	fmt.Printf("%p\n", &p) // 0xc4200160a0
	//控制数字输出的宽度和精度 %后面的数字来控制输出的宽度，默认情况下输出是右对齐的，左边加上空格
	//可以指定浮点数的输出宽度，同时还可以指定浮点数的输出精度
	// To left-justify, use the `-` flag.
	fmt.Printf("|%6d|%6d\n", 12, 345) // |    12|   345
	fmt.Printf("|%6.2f|%6.2f\n", 1.2, 3.45) // |  1.20|  3.45
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) // |1.20  |3.45  |
	fmt.Printf("|%-6.2f|%06.2f|\n", 1.2, 3.45) // |1.20  |003.45|
	//指定输出字符串的宽度来保证它们输出对齐。默认输出是右对齐的
	//左对齐：宽度之前加上 - 号
	fmt.Printf("|%6s|%6s|\n", "foo", "b") // |   foo|     b|
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b") // |foo   |b     |
	//Printf函数的输出是输出到命令行 os.Stdout 的，可以用Sprintf 来格式化后的字符串赋值给一个变量
	s := fmt.Sprintf("a %s", "string")
	fmt.Printf(s+"\n") //a string
	fmt.Fprintf(os.Stderr, "an %s\n", "error") // an error
}

