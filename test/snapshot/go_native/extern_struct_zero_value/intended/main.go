package main

import (
	"fmt"
	"image"
	"net/url"
	"snapshot/hxrt"
)

func main() {
	point := &image.Point{}
	point.X = 20
	point.Y = 22
	url := &url.URL{}
	assignedScheme := func() *string {
		url.Scheme = *hxrt.StdString(hxrt.StringFromLiteral("http"))
		return hxrt.StdString(url.Scheme)
	}()
	url.Path = *hxrt.StdString(hxrt.StringFromLiteral(""))
	url.Scheme = *hxrt.StdString(hxrt.StringConcatStringPtr(hxrt.StdString(url.Scheme), hxrt.StringFromLiteral("s")))
	appendedPath := func() *string {
		url.Path = *hxrt.StdString(hxrt.StringConcatStringPtr(hxrt.StdString(url.Path), hxrt.StringFromLiteral("/beads")))
		return hxrt.StdString(url.Path)
	}()
	if ((hxrt.StringEqualStringPtr(assignedScheme, hxrt.StringFromLiteral("http")) && hxrt.StringEqualStringPtr(hxrt.StdString(url.Scheme), hxrt.StringFromLiteral("https"))) && hxrt.StringEqualStringPtr(appendedPath, hxrt.StringFromLiteral("/beads"))) && hxrt.StringEqualStringPtr(hxrt.StdString(url.Path), hxrt.StringFromLiteral("/beads")) {
		fmt.Println(int(int32((hxrt.Int32Wrap(point.X) + hxrt.Int32Wrap(point.Y)))))
	} else {
		fmt.Println(-1)
	}
}
