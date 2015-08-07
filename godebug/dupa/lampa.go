package dupa

var a int

func init() {
	a = 1
}

func A() int {
	_ = "breakpoint"
	return a
}
