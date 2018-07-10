package main

import (
	"errors"
	"fmt"
)

type stack []int
type stackPure []int

func main() {
	var value stack
	value.push(1)
	value.push(2)
	value.push(3)
	value.push(4)
	value.push(5)
	value.pop()
	value.pop()
	value.pop()
	pop, err := value.pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pop)
	}
	fmt.Println(value)
	var valuePure stackPure
	valuePure = valuePure.push(1)
	valuePure = valuePure.push(2)
	valuePure = valuePure.push(3)
	valuePure = valuePure.push(4)
	valuePure = valuePure.push(5)
	valuePure, _, _ = valuePure.pop()
	valuePure, _, _ = valuePure.pop()
	fmt.Println(valuePure)

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
