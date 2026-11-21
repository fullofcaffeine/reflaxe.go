package main

type sys__io__FileSeek struct {
	tag    int
	params []any
}

func (self *sys__io__FileSeek) HxEnumValue() (string, []any) {
	return [3]string{"SeekBegin", "SeekCur", "SeekEnd"}[self.tag], self.params
}

var sys__io__FileSeek_SeekBegin *sys__io__FileSeek = &sys__io__FileSeek{tag: 0}

var sys__io__FileSeek_SeekCur *sys__io__FileSeek = &sys__io__FileSeek{tag: 1}

var sys__io__FileSeek_SeekEnd *sys__io__FileSeek = &sys__io__FileSeek{tag: 2}
