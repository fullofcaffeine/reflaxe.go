package main

import (
	"fmt"
	"net/http"
	"snapshot/hxrt"
)

func main() {
	result := func() *GoHTTPRequestResult {
		hx_tuple_1, hx_tuple_2 := http.NewRequest(*hxrt.StdString(hxrt.StringFromLiteral("GET")), *hxrt.StdString(hxrt.StringFromLiteral("https://example.com")), nil)
		return New_GoHTTPRequestResult(hx_tuple_1, func(err error) *go___Error {
			if err == nil {
				return nil
			}
			return New_go___Error(hxrt.StringFromLiteral(err.Error()))
		}(hx_tuple_2))
	}()
	url := result.request.URL
	fmt.Println(func() int {
		var hx_if_3 int
		if (result.error == nil) && (url != nil) {
			hx_if_3 = 42
		} else {
			hx_if_3 = -1
		}
		return hx_if_3
	}())
}
