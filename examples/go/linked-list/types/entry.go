package types

type Entry struct {
	Value    string
	Previous *Entry
	Next     *Entry
}
