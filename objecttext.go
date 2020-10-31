package objecttext

// UncheckObjectText is alias of string.
// Format check or normalize is skipped on this type.
type UncheckObjectText string

// String cast content into string.
func (v UncheckObjectText) String() (t string) {
	t = (string)(v)
	return
}

// CheckedObjectText is alias of string.
// Format check or normalize is performed on marshal and unmarshal.
type CheckedObjectText string

// String cast content into string.
func (v CheckedObjectText) String() (t string) {
	t = (string)(v)
	return
}
