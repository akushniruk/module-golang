package brackets

import "ex03"

func Bracket(str string) (bool, error) {
	stack := ex03.New()

	for count := range str {
		switch string(str[count]) {
		case "{":
			stack.Push('{')
		case "[":
			stack.Push('[')
		case "(":
			stack.Push('(')
		case "}":
			if len(stack.StackArray) != 0 {
				temp := stack.Pop()
				if temp != '{' {
					return false, nil
				}
			}
		case "]":
			if len(stack.StackArray) != 0 {
				temp := stack.Pop()
				if temp != '[' {
					return false, nil
				}
			}
		case ")":
			if len(stack.StackArray) != 0 {
				temp := stack.Pop()
				if temp != '(' {
					return false, nil
				}
			}
		default:
			return false, nil
		}
	}

	return len(stack.StackArray) == 0, nil
}
