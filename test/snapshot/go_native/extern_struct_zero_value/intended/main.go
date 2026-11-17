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
	url_local := &url.URL{}
	assignedScheme := func() *string {
		url_local.Scheme = *hxrt.StdString(hxrt.StringFromLiteral("http"))
		return hxrt.StdString(url_local.Scheme)
	}()
	url_local.Path = *hxrt.StdString(hxrt.StringFromLiteral(""))
	url_local.Scheme = *hxrt.StdString(hxrt.StringConcatStringPtr(hxrt.StdString(url_local.Scheme), hxrt.StringFromLiteral("s")))
	appendedPath := func() *string {
		url_local.Path = *hxrt.StdString(hxrt.StringConcatStringPtr(hxrt.StdString(url_local.Path), hxrt.StringFromLiteral("/beads")))
		return hxrt.StdString(url_local.Path)
	}()
	if ((hxrt.StringEqualStringPtr(assignedScheme, hxrt.StringFromLiteral("http")) && hxrt.StringEqualStringPtr(hxrt.StdString(url_local.Scheme), hxrt.StringFromLiteral("https"))) && hxrt.StringEqualStringPtr(appendedPath, hxrt.StringFromLiteral("/beads"))) && hxrt.StringEqualStringPtr(hxrt.StdString(url_local.Path), hxrt.StringFromLiteral("/beads")) {
		fmt.Println(int((hxrt.Int32Wrap(point.X) + hxrt.Int32Wrap(point.Y))))
	} else {
		fmt.Println(-1)
	}
}
