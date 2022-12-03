package collection

import "fmt"

type Collection []int

func (c Collection) Last() int {
	return c[len(c)-1]
}

func (c Collection) LastN(n int) Collection {
	return c[len(c)-n:]
}

func Test() {
	fmt.Println("testing 123")
}
