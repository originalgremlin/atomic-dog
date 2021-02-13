package scheduler

// Process TODO
type Process struct {
	// TODO proper type for process value
	processes map[ID]string
}

// Start TODO
func (p *Process) Start(image string) (ID, error) {
	return "", nil
}

// Status TODO
func (p *Process) Status(id ID) (Status, error) {
	return Unknown, nil
}

// Stop TODO
func (p *Process) Stop(id ID) error {
	return nil
}
