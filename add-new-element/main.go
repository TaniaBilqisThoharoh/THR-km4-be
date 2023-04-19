package main 
import "fmt"

func AddElement(data []int, newData int, position string) []int {
	if position == "up" {
		data = append([]int{newData}, data...)
	} else if position == "down" {
		data = append(data, newData)
	}

	return data
}

func main() {
	array := []int{1,2,3,4,5}
	fmt.Println(AddElement(array, 6, "up")) 
	fmt.Println(AddElement(array, 6, "down"))
}
