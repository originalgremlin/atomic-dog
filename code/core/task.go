package task

import (
	"fmt"

	"github.com/originalgremlin/atomic-dog/scheduler"
)

// ID TODO
type ID [16]byte

// Task TODO
type Task struct {
	Image       string
	Name        string
	Args        []interface{}
	ID          ID
	scheduler   scheduler.Scheduler
	schedulerID scheduler.ID
}

// NewTask TODO
func NewTask(image string, name string, args []interface{}, sched scheduler.Scheduler) (*Task, error) {
	var schedulerID scheduler.ID
	var err error
	if schedulerID, err = sched.Start(image); err != nil {
		return nil, err
	}
	// TODO define id as a hash of image, name, and args
	// Maybe msgpack them as an array then take a hash?
	taskID := [16]byte{}
	task := &Task{image, name, args, taskID, sched, schedulerID}
	return task, err
}

// Status TODO
func (t *Task) Status() (scheduler.Status, error) {
	return t.scheduler.Status(t.schedulerID)
}

// Stop TODO
func (t *Task) Stop() error {
	return t.scheduler.Stop(t.schedulerID)
}

// Send TODO
// TODO: if using a logical clock we need to update it for every message sent
// TODO: code smell that sending/receiving are done in two different places?
func (t *Task) Send(to ID, command string, args []interface{}) (int, error) {
	message := connection.Message{t.ID, command, args}
	if socket, ok := sockets[to]; !ok {
		return 0, fmt.Errorf("Unknown function ID: %s", id)
	} else if frames, err := message.Marshal(); err != nil {
		return 0, err
	} else {
		return socket.SendMessage(frames)
	}
}

// Broadcast a message to all tasks with a specific relation to the current task.
func (t *Task) Broadcast(to connection.Relation, command string, args ...interface{}) (int, error) {
	message := connection.Message{t.ID, command, args}
	if frames, err := message.Marshal(); err != nil {
		return 0, err
	}
	switch to {
	case connection.Parents:
		return connection.Parents.SendMessage(frames)
	case connection.Children:
		return connection.Children.SendMessage(frames)
	default:
		return 0, fmt.Errorf("Unknown relation type %s", r)
	}
}
