package main

import (
    "fmt"
    "wiselusterlab/container/stack"
)

func main() {
    s := stack.New()
    s.Swap(stack.New("Hello", 12, 3.14))
    s.Push(3 + 4i, true)

    fmt.Println(s)

    if t, e := s.Top(); e == nil {
        fmt.Printf("%d\t%#v\n", s.Size(), t)
    } else {
        panic(e)
    }

    for !s.Empty() {
        v, _ := s.Pop()
        fmt.Printf("%T\t%#v\n", v, v)
    }
}
