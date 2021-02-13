package serializer

import "github.com/vmihailenco/msgpack"

type xmsgpack struct{}

// NewMsgpack TODO
func NewMsgpack() Serializer {
	return &xmsgpack{}
}

// Serialize TODO
func (m *xmsgpack) Serialize([]interface{}) ([][]byte, error) {
	// TODO: marshal as array to utilize framing
	// return msgpack.Marshal(m)
	return nil, nil
}

// Deserialize TODO
func (m *xmsgpack) Deserialize([][]byte) ([]interface{}, error) {
	// return msgpack.Marshal(m)
	// return msgpack.Unmarshal(data, &m)
	return nil, nil
}
