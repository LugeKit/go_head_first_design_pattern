package iterator

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

type SliceIterator[T any] struct {
	slice []T
	index int
}

func NewSliceIterator[T any](slice []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		slice: slice,
	}
}

func (i *SliceIterator[T]) HasNext() bool {
	return i.index < len(i.slice)
}

func (i *SliceIterator[T]) Next() T {
	i.index += 1
	return i.slice[i.index-1]
}
