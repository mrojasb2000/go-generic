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
	/*
		for it := 0; it < len(l.data); it++ {
			if i == it {
				return l.data[it]
			}
		}*/

	for idx, value := range l.data {
		if i == idx {
			return value
		}
	}

	panic("value not found")
}
