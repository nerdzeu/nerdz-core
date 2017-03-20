package utils

// StringSet represents a string set.
type StringSet map[string]struct{}

// Contains returns true if the string is contained in the set, false otherwise.
func (ss StringSet) Contains(str string) (ok bool) {
	_, ok = ss[str]

	return
}

// Put adds the string to the set. Nothing happens if it is already present.
func (ss StringSet) Put(str string) {
	ss[str] = struct{}{}
}
