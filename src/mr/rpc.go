package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import (
	"os"
	"strconv"
)

// Enum of different job types
type JobType int

const (
	Mapper JobType = iota
	Reducer
	Waiter
)

//
// example to show how to declare the arguments
// and reply for an RPC.
//

type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y int
}

type HeartBeatAndGetTaskArgs struct {
	// NOTE: TaskID of 0 would mean not yet assigned a task
	TaskID int
}

type HeartBeatAndGetTaskReply struct {
	TaskID   int
	JobType  JobType
	Filename string
	// Ok       bool
}

type TaskReportArgs struct {
	JobType JobType
	TaskID int
}

type TaskReportReply struct {
	Ok bool
}

// Add your RPC definitions here.

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock() string {
	s := "/var/tmp/5840-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}
