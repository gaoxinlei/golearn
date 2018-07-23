//测试字符串和其他基本类型互转。
package main
import "fmt"
import "strconv"

func main() {
	var i int64 = 18
	iStr:=fmt.Sprintf("i的值是%d",i)
	fmt.Printf(iStr)
	var bo = "false"
	j,err:=strconv.ParseBool(bo)
	fmt.Printf("bo转化后的值为:%t 类型 %T err=",j,j,err)
}