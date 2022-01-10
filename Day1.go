package main

//there is nobody here
import "fmt"

var Name = "bread"

func main() {

	//var name2 = "cheese"
	//fmt.Println(subpackage.ShowData())
	//var name3 = "MUSTARD"
	//
	//var num1, num2 = 100, 300
	//num1, num2 = swap(num1, num2)
	//fmt.Println(num1, num2)
	//
	//var number1 int16
	//number1 = 6_553
	//const isPayingAttention = false
	//const pi = 3.14159
	//var number5 float64
	//number5 = 36356.3452646
	//
	//fmt.Println("test")
	//fmt.Println(Name, name2, num1, num2, pi, name3, number1, isPayingAttention, number5)
	//fmt.Println("test2")
	//fmt.Println(Raphael(3.3, 12))
}
func Raphael(grade float64, credits int) string {
	return fmt.Sprintf("Grade: %f, credits: %d", grade, credits)
}

func swap(first int, second int) (int, int) {
	return second, first
}
