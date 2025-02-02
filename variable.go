package main

import "fmt"

func main() {
	loopFunction()
}

func loopFunction() {

	var counter int16
	for i := 0; i < 10; i++ {
		counter = 0
		for j := 2; j < i%2; j++ {
			if i%j == 0 && i != 0 {
				fmt.Printf("======== %d\n", i)
				counter++
			}
		}
		if counter > 0 {
			fmt.Print("%s", i)
		}
	}
}

type Student struct {
	id      int16
	name    string
	section string
}

func variableFun() {
	var id int
	var cGPA float32
	var name, phone, section string
	status := true

	id = 102614
	cGPA = 3.9
	name = "Mamun Islam"
	phone = "01767676767"
	section = "B"
	fmt.Println(id, name, phone, section)
	fmt.Println(id, name, status)
	fmt.Printf("%.2f", cGPA)
}
