package stack

import "fmt"

type Stack struct  {
    _d []interface {}
}

func New(vals ...interface {}) *Stack {
    return &Stack{vals}
}

func (this *Stack) Empty() bool {
    return this.Size() == 0
}

func (this *Stack) Size() int {
    return len(this._d)
}

func (this *Stack) Top() (interface {}, error) {
    if this.Empty() {
        return nil, fmt.Errorf("trying to get the top of an empty stack <%p>", this)
    }

    return this._d[this.Size() - 1], nil
}

func (this *Stack) Push(vals ...interface {}) {
    this._d = append(this._d, vals...)
}

func (this *Stack) Pop() (interface {}, error) {
    top, err := this.Top()

    if err == nil {
        this._d = this._d[:this.Size() - 1]
    }

    return top, err
}

func (this *Stack) Swap(that *Stack) {
    this._d, that._d = that._d, this._d
}

func (this *Stack) String() string {
    return fmt.Sprintf("%v", this._d)
}

func (this *Stack) GoString() string {
    return "Stack" + fmt.Sprintf("%#v", this._d)[len("[]interface {}"):]
}
