package task

import (
	"fmt"

	"github.com/originalgremlin/atomic-dog/connection"
	zmq "github.com/pebbe/zmq4"
)

// initialization
var zctx *zmq.Context
var sockets map[ID]*zmq.Socket
var poller *zmq.Poller

func init() {
	zctx, _ = zmq.NewContext()
	sockets = make(map[ID]*zmq.Socket)
	poller = zmq.NewPoller()
}

// socket creation
func bind(t zmq.Type, endpoint string) (*zmq.Socket, error) {
	socket, err := zctx.NewSocket(t)
	if err == nil {
		err = socket.Bind(endpoint)
	}
	return socket, err
}

func connect(t zmq.Type, endpoint string) (*zmq.Socket, error) {
	socket, err := zctx.NewSocket(t)
	if err == nil {
		err = socket.Bind(endpoint)
	}
	return socket, err
}

// TODO: examine Tomb library
var quit chan bool = make(chan bool)

// Start the polling loop.
// TODO: if using a logical clock we need to update it for every message received
func Start() {
	for {
		select {
		case <-quit:
			return
		default:
			sockets, _ := poller.Poll(-1)
			for _, socket := range sockets {
				if bytes, err := socket.Socket.RecvMessageBytes(0); err != nil {
					fmt.Printf("Error receiving message from %s: %s\n", socket, err)
				} else {
					go handle(connection.NewMessage(bytes))
				}
			}
		}
	}
}

// Stop the polling loop and close all connections.
func Stop() {
	quit <- true
	for id, socket := range sockets {
		fmt.Printf("Closing socket %s\n", id)
		socket.Close()
		poller.RemoveBySocket(socket)
	}
	zctx.Term()
}
