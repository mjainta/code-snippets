package main

import (
	"fmt"
	"linked_list/types"
)

func main() {
	list := types.List{}

	list.Add("Hello")
	list.Add(" World")
	list.Add("!")
	list.Insert(",", 1)
	list.Insert("?", 4)
	list.Insert("# ", 0)

	for i := 0; i < len(list.Entries); i++ {
		fmt.Printf("%p %v\n", list.Entries[i], list.Entries[i])
	}

	list.PrintPretty()

	list.Remove(5)
	list.Remove(0)
	list.Remove(1)
	list.PrintPretty()
}
