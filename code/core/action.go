package task

import (
	"log"
)

// TODO: how to import Message from the parent package?
type callback func(*Message) (interface{}, error)

var callbacks map[string]callback = make(map[string]callback)

// Register TODO
func Register(name string, cb callback) error {
	callbacks[name] = cb
	return nil
}

// Unregister TODO
func Unregister(name string) error {
	delete(callbacks, name)
	return nil
}

func handle(message *Message) {
	cb, ok := callbacks[message.Command]
	if !ok {
		cb = callbacks["?"]
	}
	// TODO: inject dog global here?
	// TODO: makes it explicit what's happening, but requires a lot of boilerplate
	// TODO: maybe a better solution is to have a single global Dog namespace that gives access to Send/Broadcast/Sockets
	go func(message *Message) {
		if rv, err := cb(message); err == nil {
			// TODO: send return value back to caller as a "success" command
			sockets.Send(message.From, "success", rv)
		} else {
			// TODO: send error back to caller as an "error" command
			sockets.Send(message.From, "error", err)
		}
	}(message)
}

func init() {
	// Default actions.  Can (but likely shouldn't) be overridden by users.
	Register("connect", func(message *Message) (interface{}, error) {
		// TODO: announce self as parent or child
		return nil, nil
	})

	Register("conf", func(message *Message) (interface{}, error) {
		// TODO: provide latest configuration struct to caller
		return nil, nil
	})

	Register("ping", func(message *Message) (interface{}, error) {
		// TODO: request pong back to the caller if function is still running
		// TODO: update timestamp of last contact
		return nil, nil
	})

	Register("pong", func(message *Message) (interface{}, error) {
		// TODO: send [0.0, 1.0] completion of the task
		// TODO: update timestamp of last contact
		return nil, nil
	})

	Register("start", func(message *Message) (interface{}, error) {
		// TODO: spin up a new thread for the called function
		return nil, nil
	})

	Register("stop", func(message *Message) (interface{}, error) {
		// TODO: send quit message to the caller
		return nil, nil
	})

	Register("success", func(message *Message) (interface{}, error) {
		// TODO: spin up a new thread for the called function
		return nil, nil
	})

	Register("error", func(message *Message) (interface{}, error) {
		// TODO: spin up a new thread for the called function
		return nil, nil
	})

	// unknown command handler
	Register("?", func(message *Message) (interface{}, error) {
		log.Printf("Unknown command '%s' from %s", message.Command, message.From)
		return nil, nil
	})
}
