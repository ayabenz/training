package libs

import (
	"fmt"
)

func merge(string1, string2 string) string {
	return string2 + ", " + string1
}

func SwapStr(string1, string2 string) (string, string) {
	return string2, string1
}

func loop() {
	_ = "ini sampah"
	i := 10
	for i = 1; i <= 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}
	for ; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}
}
