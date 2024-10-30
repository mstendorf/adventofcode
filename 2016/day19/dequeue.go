package main

type Dequeue struct {
	length int
	head   *Element
	tail   *Element
}

type Element struct {
	value int
	next  *Element
	prev  *Element
}

func (s *Dequeue) Pushleft(v int) {
	if s.head == nil {
		s.head = &Element{v, nil, nil}
		s.tail = s.head
	} else {
		s.head.prev = &Element{v, s.head, nil}
		s.head = s.head.prev
	}
	s.length++
}

func (s *Dequeue) Push(v int) {
	if s.head == nil {
		s.head = &Element{v, nil, nil}
		s.tail = s.head
	} else {
		s.tail.next = &Element{v, nil, s.tail}
		s.tail = s.tail.next
	}
	s.length++
}

func (s *Dequeue) Popleft() int {
	if s.head == nil {
		return -1
	}
	res := s.head.value
	if s.head == s.tail {
		s.head = nil
		s.tail = nil
	} else {
		s.head = s.head.next
		s.head.prev = nil
	}
	s.length--
	return res
}

func (s *Dequeue) Pop() int {
	if s.head == nil {
		return -1
	}
	res := s.tail.value
	if s.head == s.tail {
		s.head = nil
		s.tail = nil
	} else {
		s.tail = s.tail.prev
		s.tail.next = nil
	}
	s.length--
	return res
}
