package main

import (
	"fmt"
	"math"
)

var gg = "good game"

// start go
func main() {
	var msg string = "Hello Pond"
	aaa := "aaa"
	var bbb = "bbb"
	// a คือ ขนม
	var a int = 100
	var b int = 200
	fmt.Println(msg)
	fmt.Println(aaa)
	fmt.Println(bbb)
	fmt.Println(a + b)
	fmt.Println(aaa + bbb)
	fmt.Println(gg)

	var price float32 = 10348.45
	var check bool = true
	fmt.Println(price)
	fmt.Println(check)

	a1, a2 := "gggg", 1111

	fmt.Println("a1 =>", a1)
	fmt.Println("a2 =>", a2)
	fmt.Printf("a2 => %#v\n : %T\n ---> type :", a2, a2)

	fmt.Printf("r : %c\n", 12222)
	fmt.Printf("r : %.2f\n", 123.336)
	fmt.Printf("ro : %#v\n", 123.336)

	var test_1 string
	var test_2 int

	fmt.Printf("test_1: %#v\n", test_1)
	fmt.Printf("test_2: %#v\n", test_2)

	num := 35

	if num%2 == 0 {
		fmt.Println("เลขคู่")
	} else {
		fmt.Println("เลขขี้")
	}

	lim := 25.6
	v := math.Pow(10, 2)

	if v < lim {
		fmt.Println("x power n is", v)
	} else {
		fmt.Printf("x power n is %g over limit %g.\n", v, lim)
	}
	var score = 49
	switch {
	case score >= 80:
		fmt.Println("A")
	case score >= 75:
		fmt.Println("B+")
	case score >= 70:
		fmt.Println("B")
	case score >= 65:
		fmt.Println("C+")
	case score >= 60:
		fmt.Println("C")
	case score >= 55:
		fmt.Println("D+")
	case score >= 50:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}

	var ttt, yyy = add(34, 56.34)
	fmt.Println("ttt =>", ttt)
	fmt.Println("yyy =>", yyy)

	var text1, text2 = sqap("pond", "karun")
	fmt.Println(text1, text2)

	var split_x, split_y = split(3464)
	fmt.Println(split_x, split_y)

	i, c := adder()
	fmt.Println(i(23))
	fmt.Println(c(45))

	var skills []string = []string{"a", "b", "c"}
	fmt.Printf("len = %#v\nvalue = %#v", len(skills), skills)

	fmt.Println(show(skills))
}

func show(arr []string) []string {
	return arr
}

func add(x float32, y float32) (float32, string) {
	return x + y, "hello"
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func sqap(x, y string) (string, string) {
	return y, x
}

func adder() (func(x int) int, func(y int) int) {
	return inc, curr
}

func inc(x int) int {
	return x + 2
}

func curr(y int) int {
	return y + 14
}
