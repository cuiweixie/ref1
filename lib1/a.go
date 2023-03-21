package lib1

import "github.com/cuiweixie/ref2/lib2"

func Lib1() string {
	return "ref2" + lib2.Lib2()
}
