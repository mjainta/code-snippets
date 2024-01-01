package types

import (
	"fmt"
)

type List struct {
	Entries []*Entry
	Head    *Entry
	Tail    *Entry
}

func (list *List) Add(value string) {
	newEntry := Entry{Value: value, Next: nil, Previous: list.Tail}
	list.Entries = append(list.Entries, &newEntry)

	if list.Head == nil {
		list.Head = &newEntry
	}

	if list.Tail != nil {
		list.Tail.Next = &newEntry
		// list.Tail = types.Entry{Value: list.Tail.Next.Value, Next: &newEntry, Previous: list.Tail}
		fmt.Println("Setting tail next to", &newEntry)
	}

	list.Tail = &newEntry
}
