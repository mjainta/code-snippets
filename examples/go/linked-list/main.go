package main

import (
	"fmt"
	"linked_list/types"
)

func main() {
	list := types.List{}

	list.Add("Hello")
	fmt.Println(list)
	list.Add("World")
	fmt.Println(list)
	list.Add("!")
	fmt.Println(&list.Entries[0], list.Entries[0])
	fmt.Println(&list.Entries[1], list.Entries[1])
	fmt.Println(&list.Entries[2], list.Entries[2])
	fmt.Printf("%p\n", list.Entries[0])
	fmt.Printf("%p\n", list.Entries[1])
	fmt.Printf("%p\n", list.Entries[2])
}
