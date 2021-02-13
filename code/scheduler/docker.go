package scheduler

// Docker TODO
type Docker struct {
	// TODO proper type for thread value
	containers map[ID]string
}

// Start TODO
func (d *Docker) Start(image string) (ID, error) {
	return "", nil
}

// Status TODO
func (d *Docker) Status(id ID) (Status, error) {
	return Unknown, nil
}

// Stop TODO
func (d *Docker) Stop(id ID) error {
	return nil
}
