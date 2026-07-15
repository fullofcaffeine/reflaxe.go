package main

type sys__io__FileSeek struct {
	tag    int
	params []any
}

var sys__io__FileSeek_SeekBegin *sys__io__FileSeek = &sys__io__FileSeek{tag: 0}

var sys__io__FileSeek_SeekCur *sys__io__FileSeek = &sys__io__FileSeek{tag: 1}

var sys__io__FileSeek_SeekEnd *sys__io__FileSeek = &sys__io__FileSeek{tag: 2}
