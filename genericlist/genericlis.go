package genericlist

type GenericList[T comparable] struct {
	data []T
}

func New[T comparable]() *GenericList[T] {
	return &GenericList[T]{
		data: []T{},
	}
}

func (l *GenericList[T]) Insert(value T) {
	l.data = append(l.data, value)
}

func (l *GenericList[T]) Get(i int) T {
	if i > len(l.data)-1 {
		panic("given index is too high")
	}

	for idx, value := range l.data {
		if i == idx {
			return value
		}
	}

	panic("value not found")
}

func (l *GenericList[T]) Remove(i int) {
	if i > len(l.data)-1 {
		panic("given index is too high")
	}

	for idx := range l.data {
		if i == idx {
			l.data = append(l.data[:idx], l.data[idx+1:]...)
		}
	}
}
