package main

type Human struct {
	name string
	age  int
	next *Human
}

type SingleLinkedList struct {
	head   *Human
	length int
}

func (sl *SingleLinkedList) Insert(name string, age int) {
	human := &Human{}
	human.name = name
	human.age = age

	if sl.length == 0 {
		sl.head = human
		sl.length++
		return
	}

	ptr := sl.head
	for i := 0; i < sl.length; i++ {
		if ptr.next == nil {
			ptr.next = human
			sl.length++
			return
		}
		ptr = ptr.next
	}
}

func (sl *SingleLinkedList) InsertAt(position int, name string, age int) {
	human := &Human{}
	human.name = name
	human.age = age

	if position < 0 {
		return
	}
	if position == 0 {
		sl.head = human
		sl.length++
		return
	}
	if position > sl.length {
		return
	}
	n := sl.GetAt(position)
	human.next = n
	prevHuman := sl.GetAt(position - 1)
	prevHuman.next = human
	sl.length++
}

func (sl *SingleLinkedList) Search(name string) *Human {
	ptr := sl.head
	for i := 0; i < sl.length; i++ {
		if ptr.name == name {
			return ptr
		}
		ptr = ptr.next
	}
	return ptr
}

func (sl *SingleLinkedList) GetAt(position int) *Human {
	ptr := sl.head
	if position < 0 {
		return ptr
	}
	if position > (sl.length - 1) {
		return nil
	}
	for i := 0; i < position; i++ {
		ptr = ptr.next
	}
	return ptr
}
