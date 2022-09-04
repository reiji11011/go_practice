package structs

import "fmt"

type Structs struct{}

func (s Structs) Test() string {
	fmt.Println("a")
	return "a"
}

func Test() string {
	fmt.Println("b")
	return "b"
}
