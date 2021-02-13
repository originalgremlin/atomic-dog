package main

import "flag"

// Arguments collects all command line flags in a single struct.
var arguments struct {
	Parent string
	ID     string
}

func init() {
	flag.StringVar(&arguments.Parent, "parent", "", "ZMQ connection address of parent process.")
	flag.StringVar(&arguments.ID, "id", "", "Unique task id for this quantum of work.")
	flag.Parse()
}
