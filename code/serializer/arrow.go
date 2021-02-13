package serializer

type xarrow struct{}

// NewArrow TODO
func NewArrow() Serializer {
	return &xarrow{}
}

// Serialize TODO
func (a *xarrow) Serialize([]interface{}) ([][]byte, error) {
	return nil, nil
}

// Deserialize TODO
func (a *xarrow) Deserialize([][]byte) ([]interface{}, error) {
	return nil, nil
}
