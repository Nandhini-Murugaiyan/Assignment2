package main
 
import (
	"fmt"
)
 
type item struct {
	value interface{} 
	next  *item
}
 
type Stack struct {
	top  *item
	size int
}
 
func (stack *Stack) Len() int {
	return stack.size
}
 
func (stack *Stack) Push(value interface{}) {
	stack.top = &item{
		value: value,
		next:  stack.top,
	}
	stack.size++
}
 
func (stack *Stack) Pop() (value interface{}) {
	if stack.Len() > 0 {
		value = stack.top.value
		stack.top = stack.top.next
		stack.size--
		return
	}
 
	return nil
}
 
func main() {
	stack := new(Stack)
	stack.Push(18)
	stack.Push(78)
	stack.Push(40)
	stack.Push(65)
	stack.Push(66)
	stack.Push(100)

	for stack.Len() > 0 {
		fmt.Println(stack.Pop())
	}
}