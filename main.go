package main

import (
	"fmt"

	"github.com/muhammadusa/tasks"
)

func main() {

	a := tasks.Fibonacci(8)
	b := tasks.Palindrome(18881)
	c, d := tasks.OddEvenSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	e := tasks.FizzBuzz(60)
	f := tasks.HasDublicate([]int{1, 2, 3, 4, 5, 6, 7, 7, 6, 8, 9, 10})

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
