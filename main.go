package main

import (
    "fmt"
    "wiselusterlab/container/stack"
)

func main() {
    s := stack.New()

    s.Swap(stack.New("Hello", 12, 3.14))
    s.Push(3 + 4i, true)

    if _, e := s.Top(); e == nil {
        fmt.Printf("%v\t%v\n", s, s.Size())
    } else {
        panic(e)
    }

    for !s.Empty() {
        v, _ := s.Pop()
        fmt.Printf("%T\t%#v\n", v, v)
    }
}
