package set

type Set struct {
	size int
}

func NewSet() *Set {
	return &Set{0}
}

func (st *Set) IsEmpty() bool {
	return st.size == 0
}

func (st *Set) Add(v string) {
	st.size = st.size + 1
}

func (st *Set) Size() int {
	return st.size
}

func (st *Set) Contains(v string) bool {
	if st.IsEmpty() {
		return false
	}
	return true
}
