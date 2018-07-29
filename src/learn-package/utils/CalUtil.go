package utils

func CalFourRegex(a int,b int) (sum int,diff int,pos int,pov int) {
	sum= a +b
	diff= a - b
	pos = a*b
	pov = a/b
	return sum,diff,pos,pov
}
