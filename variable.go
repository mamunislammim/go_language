package main

import "fmt"

func main() {
	fmt.Println("This is First line")

	studentInfo := Student{
		id:    2026,
		name:  "Mamun",
		phone: "01767676767",
		cGpa:  3.90,
	}

	fmt.Println("Student info is ", studentInfo)

	outPut := sumFun(10, 20)

	fmt.Println("Result of sum := ", outPut+50)

	fmt.Println("Student info is ", studentInfo)

	studentInfo.id = 100
	studentInfo.name = "Milon"
	fmt.Println("Student info is ", studentInfo)
	fmt.Println("Address of Pointer : ", &studentInfo)
}

type Student struct {
	id          int16
	name, phone string
	cGpa        float32
}

func sumFun(num1, num2 int16) int16 {
	result := num1 + num2
	return result
}
