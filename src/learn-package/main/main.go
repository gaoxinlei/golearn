package main

import (
	"learn-package/utils"
	"fmt"
	"strconv"
	"strings"
)

//匿名函数
var fun1 = func(x int, y int) (result int) {
	result = x * y
	return
}

//尝试闭包。
func closePack() func(int) int {
	var name = "gao"
	return func(x int) int {
		name = name + strconv.Itoa(x)
		fmt.Println("name=", name)
		return x
	}
}

//用闭包实现添加后缀名功能。
//添加后缀功能，若文件名不带闭包指定的后缀，则加上。

func addSuffix(suffix string) func (string) string {
	return func(fileName string) string {
		if(!strings.HasSuffix(fileName,suffix)){
			fileName +=suffix
		}
		return fileName
	}
}

func main() {
	sum, diff, pos, pov := utils.CalFourRegex(3, 4)
	fmt.Printf("和值%d,差值%d，积%d,商%d", sum, diff, pos, pov)
	sumFunc := getSum
	fmt.Printf("sumFunc的type：%T,getSum的type：%T", sumFunc, getSum)
	fmt.Printf("用函数参数调用getSum函数的结果:%d", sumFunc(5, 6))

	fmt.Printf("用别名调用的结果:%d", invoke(sumFunc, 5, 6))

	d := func(x, y int) int {
		return x / y
	}(5, 8)

	fmt.Printf("调用匿名函数求商:%d", d)
	d = fun1(8, 6)
	fmt.Printf("调用fun1求积:%d\n", d)

	fmt.Println("测试闭包")
	fun2 := closePack()
	fmt.Println("闭包第一次执行结果", fun2(1))
	fmt.Println("闭包第二次执行结果", fun2(2))
	fmt.Println("闭包第三次执行结果", fun2(3))

	fmt.Println("测试闭包添加.avi文件名")
	fun3 := addSuffix(".avi")
	fmt.Println("测试文件名tal100的返回值：",fun3("tal100"))
	fmt.Println("测试文件名tal101.avi的返回值：",fun3("tal101.avi"))

}

func invoke(funcParam func(int, int) int, a int, b int) (int) {
	return funcParam(a, b)
}
func getSum(a int, b int) (int) {
	return a + b
}
