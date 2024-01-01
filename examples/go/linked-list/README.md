# Go Linked List example

Did a little linked list example with Go to try out to work with structs.  
By doing that I had to get used with pointers as well as they were used in the entries list.

## How to run

`$ go run ./main.go`

It prints the list showing the pointer addresses while adding some items to the list. At the end it shows the whole list with addresses, this is where you see how the list entries are connected.

## Noticeable

* I was not able to declare a function with function body inside a struct. I aimed to use it as kind of a "class", giving it functions, but that didn't work out. One can pass a function with a body inside a struct when instantiating it, but the full body inside the struct does not seem to work. --> I believe this is so that structs focus on being a data structure, and not holding too much logic inside.
