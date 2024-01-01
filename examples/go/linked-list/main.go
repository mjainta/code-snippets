package main

import "fmt"
import "linked_list/src"
import "linked_list/types"

func main() {
	list := types.List{}
	src.Add(&list, "Hello")
	fmt.Println(list)
	src.Add(&list, "World")
	fmt.Println(list)
	src.Add(&list, "!")
	fmt.Println(&list.Entries[0], list.Entries[0])
	fmt.Println(&list.Entries[1], list.Entries[1])
	fmt.Println(&list.Entries[2], list.Entries[2])
  fmt.Printf("%p\n", list.Entries[0])
  fmt.Printf("%p\n", list.Entries[1])
	fmt.Printf("%p\n", list.Entries[2])
}