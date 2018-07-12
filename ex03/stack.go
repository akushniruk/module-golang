package stack

type Stack struct {
  stack []int
}

func New() *Stack {
  return &Stack{}
}

func (stack *Stack) Push (num int) {
	stack.stack = append(stack.stack, num)
}

func (stack *Stack) Pop() int {
  len := len(stack.stack)
  num := stack.stack[len - 1]
  stack.stack = stack.stack[:len - 1]	
  return num
}


