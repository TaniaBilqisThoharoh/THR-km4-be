package main

import (
	"fmt"
)

func removeLeftRight(arr []any, position string) []any {
	var newarray []any
	if position == "left" {
		newarray = arr[1:] 
	} else if position == "right" {
		newarray = arr[:len(arr)-1]
	}
	//yout code here!

	return newarray
}

func main() {
	fmt.Println(removeLeftRight([]any{1,2,3,4,5},"left"))
	fmt.Println(removeLeftRight([]any{1,2,3,4,5},"right"))
	fmt.Println(removeLeftRight([]any{"Edo", "Budi", "Joko", "Tono"},"left"))
	fmt.Println(removeLeftRight([]any{"Edo", "Budi", "Joko", "Tono"},"right"))
}
