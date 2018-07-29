package main

import (
	"learn-package/utils"
	"fmt"
)

func main(){
	sum,diff,pos,pov:=utils.CalFourRegex(3,4)
	fmt.Printf("和值%d,差值%d，积%d,商%d",sum,diff,pos,pov)

}
