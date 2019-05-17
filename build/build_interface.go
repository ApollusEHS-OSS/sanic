package build

import "github.com/moby/buildkit/client"

/*An Interface represents a way to output the current state of a build
  Currently there are two implementations:
  - Interactive Interfaces use advanced terminal capabilities, similar to the "curses" library
  - Plaintext Interfaces simply output jobs and their output as they are built
*/
type Interface interface {
	//StartJob notifies the interface that a job has been received, and ready to start
	StartJob(service string)
	//FailJob marks a specific job as having failed with a given error.
	//This job is "dead" and will no longer receive any logs.
	FailJob(service string, err error)
	//SucceedJob marks a specific job as having succeeded. It will no longer receive any logs.
	SucceedJob(service string)
	//ProcessStatus handles a response status from the buildkit client.
	//This could be used, for example, to process logs.
	ProcessStatus(service string, status *client.SolveStatus)
	//Terminate this interface and close any resources it is using.
	Close()
	//The interface is in charge of handling user cancelling (e.g., sigquit or ^C).
	//Call these functions when the user specifies that they would like to cancel building.
	AddCancelListener(cancelFunc func())
}
