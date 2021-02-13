package serializer

// Serializer TODO
type Serializer interface {
	Serialize([]interface{}) ([][]byte, error)
	Deserialize([][]byte) ([]interface{}, error)
}
