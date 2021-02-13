package task

// Frame TODO
type Frame []byte

// Message TODO
type Message struct {
	From    ID
	Command string
	Args    []interface{}
}

// Serialize TODO
func Serialize(m *Message) ([]Frame, error) {
	// TODO: marshal as array to utilize framing
	return nil, nil
}

// Deserialize TODO
func Deserialize(data []Frame) (*Message, error) {
	// TODO: unmarshal as array to utilize framing
	return nil, nil
}
