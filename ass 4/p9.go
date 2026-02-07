package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name  string
	Marks int
}

func main() {
	s := []Student{{"Rahul", 70}, {"Anita", 90}, {"Meena", 80}}

	sort.Slice(s, func(i, j int) bool {
		return s[i].Marks < s[j].Marks
	})

	for _, v := range s {
		fmt.Println(v.Name, v.Marks)
	}
}
