package main
import (
	"strings"
	"strconv"
	"fmt"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	newData := strings.Split(data,",")

	name:= newData[0]
	age, _:= strconv.Atoi(newData[1])
	address:= newData[2]
	return User{
		Name: name, 
		Age: age,
		Address: address,
	}
}

func main() {
	fmt.Println(ConvertData("Edo,20,Jakarta"))
	fmt.Println(ConvertData("Budi,30,Bandung"))
}
