package main

import "examples_profile_storyboard_portable/hxrt"

func main() {
	var runtime profile__StoryboardRuntime = profile__RuntimeFactory_create()
	var v any = any(Harness_render(runtime))
	hxrt.Println(v)
}
