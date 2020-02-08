package main

import (
	"fmt"
	"training/day_1/model"
)

func main() {
	//fmt.Printf("%s", merge("Hello", "World"))
	//fmt.Printf(loop())

	/*
		var names [4]string

		names[0] = "nama"
		names[1] = "saya"
		names[2] = "nggak"
		names[3] = "tahu"

		var huruf = []string{"A", "B", "C", "D", "E"}

		newHuruf := huruf[1:2]
		var aHuruf = newHuruf[1:2]
		aHuruf[0] = "Z"

		fmt.Println(huruf, len(huruf), cap(huruf))
		fmt.Println(newHuruf, len(newHuruf), cap(newHuruf))
		fmt.Println(aHuruf, len(aHuruf), cap(aHuruf))

		var string1, string2 = libs.SwapStr("World", "Hello")
		fmt.Println(string1, string2)
		fmt.Println(day_2.Test())
	*/
	/*
		var SliceMap = []map[string]int{
			{
				"A":1,
				"B":2,
			},
			{
				"A":4,
				"B":5,
			},
		}

		cFruits := append(fruits, "potato")

		fmt.Println(cFruits[4])

		var arrAngka = [2][3]string{{"1", "2", "3"}, {"4", "5", "6"}}

		for i := 0; i < len(arrAngka); i++ {
			fmt.Println(arrAngka[i])
		}

		var a = 44
		calc(&a)
	*/

	ben := model.User{ID: 1, Name: "Salamander", Dept: model.Department{ID: 2, Name: "MIS"}}
	ben2 := model.User{ID: 2, Name: "Benny", Dept: model.Department{ID: 4, Name: "MAS"}}

	fmt.Println(ben, ben2)
}

func calc(aa *int) int {
	fmt.Println(aa)
	fmt.Println(*aa)
	bb := &aa
	fmt.Println(bb)
	fmt.Println(*bb)
	return 0
}
