package linkedList

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Add(x T) {
	if l.next == nil {
		l.next = &List[T]{nil, x}
	} else {
		l.next.Add(x)
	}
}

func (l *List[T]) Print() {
	if l.next != nil {
		fmt.Print(l.val)
		l.next.Print()
	} else {
		fmt.Println(l.val)
	}
}

func Run() {
	list := List[int]{nil, 3}
	list.Add(2)
	list.Add(1)
	list.Add(0)
	list.Print()
}
