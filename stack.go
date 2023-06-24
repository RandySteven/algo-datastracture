package main

type (
	Stack struct {
		top *Plate
		len int
	}
	Plate struct {
		color string
		prev  *Plate
	}
)

func New() *Stack {
	return &Stack{nil, 0}
}

func (stack *Stack) Len() int {
	return stack.len
}

func (stack *Stack) Peek() *Plate {
	if stack.len == 0 {
		return nil
	}
	return stack.top
}

func (stack *Stack) Push(color string) {
	plate := &Plate{
		color: color,
		prev:  stack.top,
	}
	stack.top = plate
	stack.len++
}
