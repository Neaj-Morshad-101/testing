package set

type Set struct {
}

func NewSet() *Set {
	return &Set{}
}

func (st *Set) IsEmpty() bool {
	return false
}
