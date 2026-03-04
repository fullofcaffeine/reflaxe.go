package main

import "examples_profile_storyboard_portable/hxrt"

func main() {
	var runtime profile__StoryboardRuntime = profile__RuntimeFactory_create()
	hxrt.Println(Harness_render(runtime))
}
