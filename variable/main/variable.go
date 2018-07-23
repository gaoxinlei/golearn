//变量示意
package main

import "fmt"

//go语言中不允许不使用变量，以下定义的变量必须被使用。
var a int = 1
var (
	b int     = 3
	c float32 = 4.4
)

//声明时赋值，可不指定类型。
var d = 6.8
var name = "gao"

func main() {
	var position = "北京"
	var age int
	age = 28
	e := 18
	sum, sub := calNum(a, b)
	fmt.Println("a=", a, ",b=", b, ",sum=", sum, ",sub=", sub)
	fmt.Println("c=", c, ",d=", d, ",e=", e)
	fmt.Println("name=", name, ",age=", age, ",position", position)
	//汉字用整型保存字符。
	family := '高'
	fmt.Printf("family的字符值%c,编码值%d", family, family)
}

//求两个整数的和与差
func calNum(x int, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}
