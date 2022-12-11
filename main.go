package main

import (
	"education/list/storages/slice"
	"fmt"
	"os"
	"os/exec"
)

type People struct {
	ID   int
	Name string
}
type Note struct {
	name    string
	surname string
	text    string
}

func clearShell() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

type strA struct {
	name    string
	surname string
}

type strB struct {
	name string
	age  int
}

func main() {
	s := slice.NewSlice()
	_, err := s.Add(strA{"Vasya", "Ivanov"})
	if err != nil {
		fmt.Println("ERR, ", err)
	}
	_, err = s.Add(strB{"vasya", 10})
	if err != nil {
		fmt.Println("ERR, ", err)
	}

	
}
