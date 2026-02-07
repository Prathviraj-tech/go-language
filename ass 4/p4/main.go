package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	ID     int    `xml:"ID"`
	Name   string `xml:"Name"`
	Course string `xml:"Course"`
}

func main() {
	file, _ := os.Open("student.xml")
	defer file.Close()

	data, _ := ioutil.ReadAll(file)

	var s Student
	xml.Unmarshal(data, &s)

	fmt.Println(s.ID)
	fmt.Println(s.Name)
	fmt.Println(s.Course)
}
