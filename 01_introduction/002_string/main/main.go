package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var num5 int =10
	strc := strconv.Itoa(num5)
	fmt.Printf("str =%q \n",strc)

	var str string = "true"
	var b bool
	b , _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v \n",b,b)

	var xx int64 = 99999999999
	fmt.Printf("i= %v \n",xx)
	
}