package main

import (
	"fmt"
)

func main() {
	var arr[5] int = [5]int {1,2,3,4,5}
	for index,value :=range arr{
		fmt.Println("index:",index)
		fmt.Println("value:",value)
	}
	
	hero := [] string{"宋江","吴用","李逵","卢俊义","燕青"}

	fmt.Println("---------更改前-------------")
	printArr(&hero)

	modifyFirstValue(&hero)
	fmt.Println("-----------更改后-----------")
	printArr(&hero)

	//二维数组。
	var twoArr = [][] int {{1,2,},{2,5},{5,8}}
	printTwoArr(&twoArr)

	//切片。
	var slice1 [] int = [] int {1,2,3,4,5}
	fmt.Printf("slice1切片的第三号元素是:%v",slice1[3])
	slice2 := hero[1:]
	fmt.Printf("去掉第一个剩余的英雄：%v",slice2)
}
func printTwoArr(arr *[][]int) {
	for i,row:= range *arr{
		for j,col:= range row{
			fmt.Printf("行号：%v,列号：%v，值：%v\n",i,j,col)
		}
	}
}
func modifyFirstValue(arr *[]string){
	if len(*arr) >0 {
		param := *arr
		param[0]="武松"
	}
}

func printArr(arr *[]string ) {
	for index,value :=range *arr{
		fmt.Println("index:",index)
		fmt.Println("value:",value)
	}

}
