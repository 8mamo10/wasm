package main

import "fmt"

type Test struct {
	Num int
}

func (t *Test) Print() {
	fmt.Println(t)
}

func (t *Test) Twice() {
	t.Num *= 2
}

func (t *Test) Add(val int) {
	t.Num += val
}

func (t *Test) GetNum() int {
	return t.Num
}

func main() {
	var test = &Test{
		Num: 1,
	}
	test.Print()
	test.Twice()
	test.Print()
	test.Add(1)
	test.Print()
	fmt.Println(test.GetNum())
}
