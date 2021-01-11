package scheduler

// ID TODO
type ID string

// Status TODO
type Status int

// Status TODO
const (
	Starting Status = iota
	Running
	Stopping
	Stopped
	Error
	Unknown
)

// Scheduler TODO
type Scheduler interface {
	Start() (ID, error)
	Status(id ID) (Status, error)
	Stop(id ID) error
}
