package hxrt

import "fmt"

func Println(value any) {
	fmt.Println(*StdString(value))
}
