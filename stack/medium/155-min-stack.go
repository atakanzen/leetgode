type MinStack struct {
    Stack []int
    Mins []int
}


func Constructor() MinStack {
    return MinStack{
        Stack: make([]int, 0),
        Mins: make([]int, 0),
    }
}


func (this *MinStack) Push(val int)  {
    this.Stack = append(this.Stack, val)
    if len(this.Mins) == 0 ||  val < this.Mins[len(this.Mins) - 1] {
        this.Mins = append(this.Mins, val)
    } else {
        this.Mins = append(this.Mins, this.Mins[len(this.Mins)- 1])
    }
}


func (this *MinStack) Pop()  {
    this.Stack = this.Stack[:len(this.Stack) - 1]
    this.Mins = this.Mins[:len(this.Mins) - 1]
}


func (this *MinStack) Top() int {
    return this.Stack[len(this.Stack)-1]
}


func (this *MinStack) GetMin() int {
    if len(this.Mins) > 0 {
        return this.Mins[len(this.Mins) - 1]
    }

    return 0
}
