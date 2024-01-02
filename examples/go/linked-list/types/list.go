package types

import (
	"fmt"
)

type List struct {
	Entries []*Entry
	Head    *Entry
	Tail    *Entry
}

func (list *List) PrintPretty() {
	for i := 0; i < len(list.Entries); i++ {
		fmt.Printf("%s", list.Entries[i].Value)
		if i == len(list.Entries)-1 {
			fmt.Printf("\n")
		}
	}
}

func (list *List) Add(value string) {
	list.Insert(value, len(list.Entries))
}

func (list *List) Insert(value string, index int) {
	if len(list.Entries) <= index {
		newEntry := Entry{Value: value, Next: nil, Previous: list.Tail}
		list.Entries = append(list.Entries, &newEntry)

		if list.Head == nil {
			list.Head = &newEntry
		}

		if list.Tail != nil {
			list.Tail.Next = &newEntry
			fmt.Println("Setting tail next to", &newEntry)
		}

		list.Tail = &newEntry
	} else if index == 0 {
		newEntry := Entry{Value: value, Next: list.Entries[0], Previous: nil}
		list.Entries = append([]*Entry{&newEntry}, list.Entries...)
		list.Entries[1].Previous = &newEntry
	} else {
		newEntry := Entry{}
		if len(list.Entries) > index {
			newEntry = Entry{Value: value, Next: list.Entries[index], Previous: list.Entries[index-1]}
			list.Entries = append(list.Entries[:index+1], list.Entries[index:]...)
		} else {
			newEntry = Entry{Value: value, Next: nil, Previous: list.Entries[index-1]}
			list.Entries = append(list.Entries, &newEntry)
		}
		list.Entries[index] = &newEntry
		list.Entries[index-1].Next = &newEntry

		if len(list.Entries) > index+1 {
			list.Entries[index+1].Previous = &newEntry
		}

		if list.Head == nil {
			list.Head = &newEntry
		}
	}
}

func (list *List) Remove(index int) {
	if index == 0 {
		list.Entries = list.Entries[1:]
		list.Entries[0].Previous = nil
	} else if index == len(list.Entries)-1 {
		list.Entries = list.Entries[:len(list.Entries)-1]
		list.Entries[len(list.Entries)-1].Next = nil
	} else {
		list.Entries = append(list.Entries[:index], list.Entries[index+1:]...)
		list.Entries[index].Previous.Next = list.Entries[index]
		list.Entries[index].Next.Previous = list.Entries[index]
	}
}
