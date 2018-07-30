package main

import (
	"fmt"
	"time"
	"errors"
)

func throwException()  {
	var a int = 10
	var b int = 0
	//处理异常。
	defer func(){
		if err:=recover();err !=nil{
			fmt.Println("收到错误消息:",err)
		}
	}()
	var c int = a/b
	fmt.Println("商值：",c)
}

func errorCondtion(param int) {
	var err error
	if 2==param {
		err = errors.New("为2")
	}
	if err!=nil {
		panic(err)
	}else{
		fmt.Println("不满足抛异常条件")
	}
}

//测试defer延迟执行。
func main(){
	start:= time.Now().Second()
	sum :=getSum(5,6)
	end := time.Now().Second()
	fmt.Println("getSume结果:",sum)//4
	fmt.Println("getSum耗时:",end-start)

	fmt.Println("调用抛异常方法")
	throwException()
	fmt.Println("主程序正常向下执行。")
	errorCondtion(2)
}

func getSum(x int,y int) (sum int){
	defer  fmt.Printf("x=",x)//3
	defer fmt.Printf("y=",y)//2
	sum = x +y
	x++
	y++
	fmt.Println("sum=",sum)//1.
	return
}
