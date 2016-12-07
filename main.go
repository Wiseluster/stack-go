package main

import (
    "fmt"
    "reflect"
    "wiselusterlab/container/stack"
)

func main() {
    s := stack.New("Hello", 12, 3.14)
    s.Push(3 + 4i)

    if t, e := s.Top(); e == nil {
        fmt.Println(s.Size(), t)
    } else {
        panic(e)
    }

    for !s.Empty() {
        v, _ := s.Pop()
        fmt.Println(reflect.TypeOf(v), v)
    }
}
