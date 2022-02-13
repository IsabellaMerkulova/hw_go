package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	length    int
	firstElem *ListItem
	lastElem  *ListItem
}

func (l list) Len() int {
	return l.length
}

func (l list) Front() *ListItem {
	return l.firstElem
}

func (l list) Back() *ListItem {
	return l.lastElem
}

func (l *list) PushFront(v interface{}) *ListItem {
	l.length++
	newElem := ListItem{v, l.firstElem, nil}
	if l.firstElem != nil {
		l.firstElem.Prev = &newElem
	} else {
		l.lastElem = &newElem
	}
	l.firstElem = &newElem
	return &newElem
}

func (l *list) PushBack(v interface{}) *ListItem {
	l.length++
	newElem := ListItem{v, nil, l.lastElem}
	if l.lastElem != nil {
		l.lastElem.Next = &newElem
	} else {
		l.firstElem = &newElem
	}
	l.lastElem = &newElem
	return &newElem
}

func (l *list) Remove(i *ListItem) {
	l.length--
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.firstElem = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.lastElem = i.Prev
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.firstElem {
		return
	}
	i.Prev.Next = i.Next
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	i.Prev = nil
	i.Next = l.firstElem
	l.firstElem.Prev = i
	l.firstElem = i
}

func NewList() List {
	return new(list)
}
