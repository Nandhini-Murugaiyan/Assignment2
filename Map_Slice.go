package main

import "fmt"

func main() {
	fmt.Println(make(map[string]int)) 
	m := make(map[string]int, 3)
	fmt.Println(m, len(m)) 
	m["C"] = 1972
	m["Go"] = 2009
	fmt.Println(m, len(m)) 
	s := make([]int, 3, 5)
	fmt.Println(s, len(s), cap(s)) 
	s = make([]int, 2)
	fmt.Println(s, len(s), cap(s)) 
}