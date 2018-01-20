package main 
import( 
	"fmt" //
	"unsafe"
	// "math"
	// "time"
)



func getAverage(arr []int, size int) float32 {
	var i int
	var sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum / size)

	return avg;
}


/*
几种函数：
1. man 没有返回值的函数
2. 普通函数
3. 函数返回多个值
4. 闭包函数
5. 递归函数
6. 类型方法，类似C++中类的成员函数
7. 接口和多肽
8. 错误处理，Defer接口
9. 错误处理，Panic/Recover
*/


/*
函数思想：
	可以通过函数划分不同功能，逻辑上每个函数执行的是指定的任务。
	函数声明告诉了编译器函数的名称，返回类型和参数。

函数参数：
	值传递：调用函数时，将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
	引用传递：在调用函数时，将实际参数的地址传递到函数中，在函数中对参数所进行的修改，将影响到实际参数。
go 语言默认使用的是值传递。

函数的用法：
	函数作为值。
	闭包：匿名函数，可在动态编程中使用。
	方法：一个包含了接受者的函数。

*/


/*
go语言变量作用域：已经声明的标识符所表示的常量、类型、变量、函数或包在源代码中的作用域范围。
变量可以声明在三个地方：
	局部变量：函数内定义的变量。
	全局变量：函数外定义的变量。-可以在整个包甚至外部包（被导出后）使用。
	形式参数：函数定义中的变量。
	go 中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。

	初始化局部和全局变量默认值： 
		int -> 0 
		float32 -> 0
		pointer -> nil
*/



//值类型和引用类型
/*
值类型:
	 int、float、bool 和 string 
	 使用这些类型的变量直接指向存在内存中的值.
	 当使用等号 = 将一个变量的值赋值给另一个变量时，如：j = i，实际上是在内存中将 i 的值进行了拷贝。

	 你可以通过 &i 来获取变量 i 的内存地址，例如：0xf840000040（每次的地址都可能不一样）。值类型的变量的值存储在栈中。
内存地址会根据机器的不同而有所不同，甚至相同的程序在不同的机器上执行后也会有不同的内存地址。因为每台机器可能有不同的存储器布局，并且位置分配也可能不同。
更复杂的数据通常会需要使用多个字，这些数据一般使用引用类型保存。
*/ 



var a = "Go-Programer"
var b string = "go.com"
var c bool


var x, y int
var (  // 这种因式分解关键字的写法一般用于声明全局变量
	num1 int
	numBool1 bool
)

var num2, num3 int = 1, 3
var num4, num5 = 123, "hello"

var f float32 = 1.6
var s = "abc"



func main(){
	
	// w := 123
	// println(w)



	// var i = 3
	// go func (a int)  {
	// 	println(a)
	// 	println(1)
	// }(i)
	// println(2)
	// // time.Sleep(1 * time.Second)


	// ch := make(chan int)
	// ch <- 1
	// go func ()  {
	// 	<- ch
	// 	println("1")
	// }()
	// println("2")
	// // fatal error: all goroutines are asleep - deadlock!
	// // 线程陷入死锁。


	s := []int{7, 2, 8, 9}
	
	// c := make(chan int)
	var c chan int
	go sum(s[:len(s)/2], c)
	// go sum(s[len(s)/2:], c)
	x := <-c
	// x, y := <-c, <-c 



	// close(c)
	// z := <-c
	// go func (chan int)  {
	// 	z := 3
	// 	c <- z
	// 	fmt.Println(z)	
	// }(c)
	


	// z := 3
	// c <- z
	// fmt.Println(z)
		

	fmt.Println(c)

	fmt.Println(x, y, x+y)




	// 从一个被 close 的 channel 中接收数据，会立即返回元素类型的零值。
	// 从一个 nil 的 channel 中接收数据，会发生死锁。  fatal error: all goroutines are asleep - deadlock!
	// 向一个被 close 的 channel 中继续发送数据，会导致  panic: send on closed channel
	// 一个没有初始化的 channel 是 nil ，nil 的 channel receive 或 send 会导致死锁。 fatal error: all goroutines are asleep - deadlock!
}


func sum(s []int, c chan int)  {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}



func demo() {
	// fmt.Println(Max(num3, num2))

	fmt.Println("Hello, World!")
	println(x, y, a, b, c, num1, num2, num3, num4, num5)


	//这种不带声明格式的只能在函数体中出现
	g, h := 123, "hello"
	println(g, h)

	// var aString = "abcd"

	aVar, bVar, cVar := 2, 4, "ccc" 
	println(aVar, bVar, cVar)




	// Go 语言常量
	// const identifier [type] = value
	// 可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型。
	const LENGTH int = 10
	const WIDTH int = 3
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Println("面积为：%d", area)
	println()
	println(a, b, c)

	
	const (
		Unknown = 0
		Female = 1
		Male = 2
	)

	// 常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。
	// 常量表达式中，函数必须是内置函数，否则编译不过：
	
	const (
		constA = "abc"
		constB = len(constA)
		constC = unsafe.Sizeof(constA)
	)
	
	
	println(constA, constB, constC)

	
	// iota , 特殊常量， 可以认为是一个可以被编译器修改的常量。
	// 在每一个const关键字出现时，被重置为0， 然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1。
	// iota 可以被用作枚举类型
	const (
		iotaA = iota
		iotaB = iota
		iotaC = iota
	)
	// 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：
	const (
		iotaAA = iota
		iotaBB
		iotaCC
	)

	const (
		iotaAAA = iota //0
		iotaBBB //1
		iotaCCC //2
		iotaDDD = "hello"  //"hello" 独立值， iota += 1  
		iotaEEE  //"hello"  iota += 1
		iotaFFF = 100  //100
		iotaGGG   //100 iota += 1
		iotaHHH = iota //7 , 恢复计数
		iotaIII //8 
	)
	fmt.Println(iotaAAA, iotaBBB, iotaCC, iotaDDD, iotaEEE, iotaFFF, iotaGGG, iotaHHH, iotaIII)


	const (
		i = 1<<iota
		j = 3<<iota
		k
		l
	)

	fmt.Println("i=",i)  // 1
	fmt.Println("j=",j)  // 6
	fmt.Println("k=",k)  // 12
	fmt.Println("l=",l)  // 24

	//iota表示从0开始自动加1，所以 i=1<<0, j=3<<1 (<<表示左移)，即：i=1,j=6,这个问题，关键在k和l，从输出结果看：k=3<<2, l=3<<3 ]



	//算术运算符的用法：
	var cN1 = 21
	var cN2 = 10
	var cN3 int

	cN3 = cN1 + cN2 
	fmt.Println("第一行 - cN3 的值为 %d \n", cN3)

	var cN4 = cN1 - cN2
	fmt.Println("第二行 - cN4 的值为 %d \n", cN4)

	var cN5 = cN1 * cN2
	fmt.Println("第三行 - cN5 的值为 %d \n", cN5) 

	fmt.Println(cN1 % cN2)
	fmt.Println(cN1 / cN2)
	cN1++
	fmt.Println(cN1)
	cN2--
	fmt.Println(cN2)
	

	fmt.Println(cN1 == cN2) //false
	fmt.Println(cN1 != cN2) //true
	fmt.Println(cN1 > cN2) //true
	fmt.Println(cN1 < cN2) //false
	fmt.Println(cN1 >= cN2) //true
	fmt.Println(cN1 <= cN2)//false


	if cN1 == cN2 {
		fmt.Println("相等")
	}else{
		fmt.Println("不相等")
	}

	var aBool = true
	var bBool = false
	if aBool && bBool {
		fmt.Println("true 与运算")
	}else{
		fmt.Println("false 与运算")
	}

	if aBool || bBool {
		fmt.Println("true 或运算")
	}else{
		fmt.Println("false 或运算")
	}

	if !aBool {
		fmt.Println("! 非运算 ")
	}else{
		fmt.Println("! !!!非运算 ")
	}


	//Switch
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90: grade = "A"
	case 80: grade = "B"
	case 50, 60, 70 : grade = "C"
	default: grade = "D"

	}

	fmt.Println("等级： %s\n", grade)

	//Type Switch
	// switch 语句还可以被用于type-switch来判断某个interface变量中实际存储的变量类型。
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Println(" x 的类型 ： %T", i)
	case int:
		fmt.Println(" x 是 int 型")
	case float64:
		fmt.Println(" x 是float64 型")
	case func(int) float64 :
		fmt.Println("x 是 func(int) 型")
	case bool, string:
		fmt.Println(" x 是bool 或 string 型")
	default:
		fmt.Println("未知类型")
	}


	//循环语句
	var cj int = 15
	var ck int

	for ck := 0; ck < 10; ck++ {
		fmt.Println("ck 的值为： %d\n", ck)
	}
	
	
	for ck < cj {
		ck ++
		fmt.Println("ck++ 后的值为： %d\n", ck)
	}


	numbers := [6]int{1, 2, 3, 5}
	for key, value := range numbers {
		fmt.Println("第 %d 位的值：%d\n",key , value)
	}
	

	//函数定义
	/*
	func function_name([parameter list]) [return_types] {
		//函数体
	}
	函数定义解析：
		func: 函数有func开始声明。
		function_name: 函数名，函数名和参数列表一起构成了函数签名。
		parameter list: 参数列表，参数就像一个占位符，当函数被调用时，可以将值传递给参数，这个值被称为实际参数。
			参数列表指定的是参数类型、顺序、及参数个数。参数是可以选的，也就是说函数也可以不包含参数。
		return_types: 返回类型，函数返回一列值。return_types
	*/

	calculate()



/*
数组：
	go数组声明需要指定元素个数，语法格式如下：

	
*/ 
	var balance = [5]float32{1,2,3,4,5}
	var balanceT = []int{4,5,6,7,8}

	fmt.Println(balance)
	fmt.Println(balanceT)

	balanceT[0] = 1
	fmt.Println(balanceT)


	var avg float32
	avg = getAverage(balanceT, 5)
	fmt.Println("平均值：%f",avg)


	fmt.Println("变量地址：%x",&avg)
	

}





func calculate(){
	// var a int = 100
	// var b int = 200
	// var ret int

	
	// ret = Max(a, b)
	// fmt.Println("最大值为：%d\n",ret)
}












