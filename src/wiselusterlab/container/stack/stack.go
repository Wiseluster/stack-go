package stack

import "fmt"

type Stack struct  {
    data []interface {}
}

func New(vals ...interface {}) *Stack {
    return &Stack{vals}
}

func (this *Stack) Empty() bool {
    return this.Size() == 0
}

func (this *Stack) Size() int {
    return len(this.data)
}

func (this *Stack) Top() (interface {}, error) {
    if this.Empty() {
        return nil, fmt.Errorf("trying to get the top of an empty stack <%p>", this)
    }

    return this.data[this.Size() - 1], nil
}

func (this *Stack) Push(vals ...interface {}) {
    this.data = append(this.data, vals...)
}

func (this *Stack) Pop() (interface {}, error) {
    top, err := this.Top()

    if err == nil {
        this.data = this.data[:this.Size() - 1]
    }

    return top, err
}

func (this *Stack) Swap(that *Stack) {
    this.data, that.data = that.data, this.data
}

func (this *Stack) String() string {
    return "Stack" + fmt.Sprintf("%#+v", this.data)[len("[]interface {}"):]
}
