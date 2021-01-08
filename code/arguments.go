package main

import "flag"

// Arguments collects all command line flags in a single struct.
var Arguments struct {
	Parent string
	ID     string
}

func init() {
	flag.StringVar(&Arguments.Parent, "parent", "", "ZMQ connection address of parent process.")
	flag.StringVar(&Arguments.ID, "id", "", "Unique task id for this quantum of work.")
	flag.Parse()
}
