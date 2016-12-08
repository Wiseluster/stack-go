package main

import (
    "fmt"
    "wiselusterlab/container/stack"
)

type Foo struct {
    Bar int64
}

func main() {
    s := stack.New()

    s.Swap(stack.New("Hello", 12, 3.14))
    s.Push(3 + 4i, true, Foo{9})

    fmt.Println(s)

    if _, e := s.Top(); e == nil {
        fmt.Printf("%#v\t%v\n", s, s.Size())
    } else {
        panic(e)
    }

    for !s.Empty() {
        v, _ := s.Pop()
        fmt.Printf("%T\t%#[1]v\n", v)
    }
}
