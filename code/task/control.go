package task

import (
	"crypto/rand"
	"encoding/hex"
)

// ID Globally unique identifier for this node and it's relations.
// XXX: do we need a node id?  For logging?
type ID [16]byte

// NewID TODO
func NewID() (ID, error) {
	id := make([]byte, 16)
	_, err := rand.Read(id)
	return id, err
}

// Hex TODO
func (id *ID) String() string {
	return hex.EncodeToString(id)
}

var self ID = NewID()

// Relation between this node and the connection.
type Relation int

// Relation between this node and the connection.
const (
	Parents Relation = iota
	Children
)
