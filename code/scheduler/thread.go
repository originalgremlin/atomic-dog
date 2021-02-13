package scheduler

// Thread TODO
type Thread struct {
	// TODO proper type for thread value
	threads map[ID]string
}

// Start TODO
func (t *Thread) Start(image string) (ID, error) {
	return "", nil
}

// Status TODO
func (t *Thread) Status(id ID) (Status, error) {
	return Unknown, nil
}

// Stop TODO
func (t *Thread) Stop(id ID) error {
	return nil
}
