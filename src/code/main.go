package main

import (
	"container/list"
	"fmt"
)

func main() {
	str := "({{}})[][](({{}}))"
	f := true
	l := list.New()
	for i := 0; i < len(str); i++ {
		if str[i] == '[' {
			l.PushBack(1)
		}
		if str[i] == ']' {
			if l.Back().Value != 1 {
				f = false
				break
			} else {
				l.Remove(l.Back())
			}
		}
		if str[i] == '(' {
			l.PushBack(2)
		}
		if str[i] == ')' {
			if l.Back().Value != 2 {
				f = false
				break
			} else {
				l.Remove(l.Back())
			}
		}
		if str[i] == '{' {
			l.PushBack(3)
		}
		if str[i] == '}' {
			if l.Back().Value != 3 {
				f = false
				break
			} else {
				l.Remove(l.Back())
			}
		}
	}

	fmt.Println(f)
}
