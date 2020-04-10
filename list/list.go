package list

import (
	"errors"
)

type node struct {
	next *node
	data interface{}
}

type List struct {
	// Attributes
	fist  *node
	last  *node
	count int
}

func (l *List) IsEmpty() bool {
	return l.fist == nil
}

func (l *List) ItemsCount() int {
	return l.count
}

func (l *List) AddElement(data interface{}) error {
	if l.fist == nil {
		temp := node{next: nil, data: data}
		l.fist = &temp
		l.last = l.fist
		l.count++
	} else {
		temp := node{next: nil, data: data}
		l.last.next = &temp
		l.last = &temp
		l.count++
	}
	return nil
}

func (l *List) GetElement(item int) (interface{}, error) {
	if item > l.count || item < 0 {
		return nil, errors.New("Invalid item number")
	}
	i := 0
	var data interface{}
	temp := l.fist
	for {
		if i == item {
			data = temp.data
			break
		}
		temp = temp.next
		i++
	}
	return data, nil
}

func (l *List) UpdateElement(item int, newData interface{}) error {
	if item > l.count || item < 0 {
		return errors.New("Invalid item number")
	}
	i := 0
	temp := l.fist
	for {
		if i == item {
			temp.data = newData
			break
		}
		temp = temp.next
		i++
	}
	return nil
}

func (l *List) RemoveElement(item int) error {
	if item > l.count || item < 0 {
		return errors.New("Invalid item number")
	}
	// Remove first element
	if item == 0 {
		temp := l.fist
		l.fist = l.fist.next
		temp.next = nil
		temp.data = nil
		l.count--

		return nil
	}
	// Remove element
	i := 0
	temp := l.fist
	for {
		temp0 := temp.next
		i++
		if i == item {
			if temp0.next == nil {
				// temp0 is the las item
				l.last = temp
				l.last.next = nil
				temp0.data = nil
			} else {
				// temp0 is not the las item
				temp.next = temp0.next
				temp0.next = nil
				temp0.data = nil
			}
			l.count--
			break
		}
		temp = temp.next
	}
	return nil
}
