package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

// type definitions

// ID TODO foo
type ID string

// Relation TODO
type Relation int

// Relation TODO
const (
	Parent Relation = iota
	Child
)

// initialization
var zctx *zmq.Context
var sockets map[ID]*zmq.Socket
var poller *zmq.Poller

func init() {
	zctx, _ = zmq.NewContext()
	sockets = make(map[ID]*zmq.Socket)
	poller = zmq.NewPoller()
	// TODO: open all sockets
}

func handle(socket *zmq.Socket, message [][]byte) {

}

// socket creation
func subscriber(addr string, topics []string) (*zmq.Socket, error) {
	socket, err := zctx.NewSocket(zmq.SUB)
	if err == nil {
		socket.Connect(addr)
		poller.Add(socket, zmq.POLLIN)
		for _, topic := range topics {
			socket.SetSubscribe(topic)
		}
	}
	return socket, err
}

func publisher(addr string) (*zmq.Socket, error) {
	socket, err := zctx.NewSocket(zmq.PUB)
	if err == nil {
		socket.Bind(addr)
	}
	return socket, err
}

// TODO: bind router or dealer?
func router(addr string) (*zmq.Socket, error) {
	socket, err := zctx.NewSocket(zmq.ROUTER)
	if err == nil {
		socket.Connect(addr)
		poller.Add(socket, zmq.POLLIN)
	}
	return socket, err
}

func dealer(addr string) (*zmq.Socket, error) {
	socket, err := zctx.NewSocket(zmq.DEALER)
	if err == nil {
		socket.Bind(addr)
		poller.Add(socket, zmq.POLLIN)
	}
	return socket, err
}

// socket control

// Start TODO
// TODO: control channel to break out of the loop
func Start() {
	for {
		sockets, _ := poller.Poll(-1)
		for _, socket := range sockets {
			if message, err := socket.Socket.RecvMessageBytes(0); err != nil {
				fmt.Printf("Error receiving message from %s", socket)
			} else {
				go handle(socket.Socket, message)
			}
		}
	}
}

// Stop TODO
func Stop() {
	for id, socket := range sockets {
		fmt.Printf("Closing socket %s\n", id)
		poller.RemoveBySocket(socket)
		socket.Close()
	}
	zctx.Term()
}

// Send TODO
func Send(node ID, message [][]byte) (int, error) {
	if socket, ok := sockets[node]; ok {
		return socket.SendMessage(message)
	} else {
		return 0, fmt.Errorf("Unknown node ID: %s", node)
	}
}

// Broadcast TODO
func Broadcast(relation Relation, topic string, message [][]byte) (int, error) {
	switch relation {
	case Parent:
		return sockets["ToParents"].SendMessage(topic, message)
	case Child:
		return sockets["ToChildren"].SendMessage(topic, message)
	default:
		return 0, fmt.Errorf("Unknown relation type %s", relation)
	}
}
