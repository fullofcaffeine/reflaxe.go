package hxrt

import "fmt"

func StringFromLiteral(value string) *string {
	copy := value
	return &copy
}

func Println(value *string) {
	if value == nil {
		fmt.Println("null")
		return
	}
	fmt.Println(*value)
}
