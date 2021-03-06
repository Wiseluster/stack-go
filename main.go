package main

import (
    "fmt"
    "strconv"
    "wiselusterlab/container/stack"
)

type char rune

func (this char) String() string {
    return string(this)
}

func (this char) GoString() string {
    return strconv.QuoteRune(rune(this))
}

func main() {
    s := stack.New()

    s.Swap(stack.New("Hello", 12, 3.14))
    s.Push(3 + 4i, true, char('A'))

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
