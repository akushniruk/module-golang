package ex03

type Stack struct {
	StackArray []int
}

func New() *Stack {
	return &Stack{}
}

func (stack *Stack) Push(num int) {
	stack.StackArray = append(stack.StackArray, num)
}

func (stack *Stack) Pop() int {
	len := len(stack.StackArray)
	num := stack.StackArray[len-1]
	stack.StackArray = stack.StackArray[:len-1]
	return num
}
