package main

import "fmt"

type Stack []string

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *Stack) Pop() string {
	if len(*s) == 0 {
		return ""
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack) Peek() string {
	if len(*s) == 0 {
		return ""
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) String() string {
	return fmt.Sprintf("%v", []string(*s))
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func main() {
	var s Stack
	s.Push("a")
	s.Push("b")
	s.Push("c")
	fmt.Println(s.String())
	fmt.Println(s.Pop())
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Peek())
}
