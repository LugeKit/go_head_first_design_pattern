package iterator

type SliceMenu struct {
	items []int
}

func (m *SliceMenu) Iterator() Iterator[int] {
	return NewSliceIterator(m.items)
}
