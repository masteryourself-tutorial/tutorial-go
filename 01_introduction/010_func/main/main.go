package main

import (
	"fmt"
)

func main()  {
	calc(3.2, 6.8, '+')
	calc(1.1, 1.1, '*')
}

func calc(num1 float64, num2 float64, operation byte) (float64) {
	var res float64
	switch operation {
		case '+':
			res = num1 + num2
		case '-':
			res = num1 - num2
		case '*':
			res = num1 * num2
		case '/':
			res = num1 / num2
		default :
			fmt.Println("操作符号错误", operation)
	}
	fmt.Printf("%v %c %v = %v \n", num1, operation, num2, res)
	return res

}