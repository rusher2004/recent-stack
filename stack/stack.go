package stack

type Item[T any] struct {
	ID   string
	Type string
	Data T
}

type Stack[T any] struct {
	items []Item[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]Item[T], 0),
	}
}

func (s *Stack[T]) Push(item Item[T]) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (Item[T], bool) {
	if len(s.items) == 0 {
		var zero Item[T]
		return zero, false
	}
	
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	
	return item, true
}

func (s *Stack[T]) Peek() (Item[T], bool) {
	if len(s.items) == 0 {
		var zero Item[T]
		return zero, false
	}
	
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

func (s *Stack[T]) Clear() {
	s.items = make([]Item[T], 0)
}

func (s *Stack[T]) Items() []Item[T] {
	return s.items
}