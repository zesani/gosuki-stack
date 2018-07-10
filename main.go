package main

import (
	"errors"
	"fmt"
)

type stack []int
type stackPure []int

func main() {
	var stack stack
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)
	stack.push(5)
	stack.pop()
	stack.pop()
	stack.pop()
	pop, err := stack.pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pop)
	}
	fmt.Println(stack)
	var stackPure stackPure
	stackPure = stackPure.push(1)
	stackPure = stackPure.push(2)
	stackPure = stackPure.push(3)
	stackPure = stackPure.push(4)
	stackPure = stackPure.push(5)
	stackPure, _, _ = stackPure.pop()
	stackPure, _, _ = stackPure.pop()
	fmt.Println(stackPure)

}

func (s *stack) push(number int) {
	*s = append(*s, number)
}

func (s *stack) pop() (int, error) {
	temp := *s
	if len(temp) == 0 {
		return 0, errors.New("Can't pop stack")
	}
	pop := temp[len(temp)-1]
	*s = append(temp[:len(temp)-1], temp[len(temp)-1+1:]...)
	return pop, nil
}

func (s stackPure) push(number int) stackPure {
	return append(s, number)
}

func (s stackPure) pop() (stackPure, int, error) {
	if len(s) == 0 {
		return s, 0, errors.New("Can't pop stack")
	}
	pop := s[len(s)-1]
	s = append(s[:len(s)-1], s[len(s)-1+1:]...)
	return s, pop, nil
}
