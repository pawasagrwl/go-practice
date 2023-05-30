package main

import ("fmt")

func main() {
	var a int = 10
	var b bool = true
	var str string = "Hello World!"
	var f float32 = 1.234
	fmt.Printf("a = %d, b = %t, str = %s, f = %f\n", a, b, str, f)
	fmt.Printf("a = %v, b = %v, str = %v, f = %v\n", a, b, str, f)
	fmt.Printf("a = %T, b = %T, str = %T, f = %T\n", a, b, str, f)
	fmt.Printf("a = %p, b = %p, str = %p, f = %p\n", &a, &b, &str, &f)
	fmt.Printf("a = %#v, b = %#v, str = %#v, f = %#v\n", a, b, str, f)
	var input int
	fmt.Scanf("%d", &input)
	fmt.Println("input =", input)
	if (a > 0) {
		fmt.Println("a > 0")
	} else {
		fmt.Println("a <= 0")
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
