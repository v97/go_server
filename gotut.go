package main

import ("fmt")

type car struct {
	acc_ped float64
	break_ped float64
	steer float64
}

func guess_kph(c car) float64 {
	return 102.0
}

func add(x float64, y float64) float64 {
	return x + y
}

func foo(){
	num1 := 4.9
	num2 := 1.9

	sumnum := add(num1, num2)
	fmt.Println(sumnum)
}

func stringTest(a, b, c string) (string, string, string) {
	return c,b,a
}

func main() {
	foo()
	fmt.Println(stringTest("you", "love", "I"))

	a := 3.0
	b := &a

	fmt.Println(a, b, *b)

	acar := car{100, 0, 100}

	fmt.Println(acar.acc_ped, guess_kph(acar))
}