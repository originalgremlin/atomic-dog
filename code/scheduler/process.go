package scheduler

// Local TODO
type Local struct {
	// TODO proper type for process value
	processes map[ID]string
}

// Start TODO
func (l *Local) Start() (ID, error) {
	return "", nil
}

// Status TODO
func (l *Local) Status(id ID) (Status, error) {
	return Unknown, nil
}

// Stop TODO
func (l *Local) Stop(id ID) error {
	return nil
}
