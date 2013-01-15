package ovg

import (
	"net"
	"net/rpc"
)

type ProjectNotifier interface {
	/*
	   This call indicates that a particular worker has become unresponsive. The arguments passed are :-
	       * deadWorker- The id of the worker who has gone dead.
	       * replaced - this indicates wether the dead worker could be replaced with another one or not. If the worker is replaced then
	           it retains the same id i.e. deadWorker
	*/
	WorkerDead(deadWorker int, replaced bool)

	/*
	   This call signals that a new worker is available for this project and should be added to the list of workers. The single argument
	   newWorker is the id assigned to this worker. This id can be one previously assigned to a dead worker that could not be replaced.
	*/
	NewWorker(newWorker int)
}

// The type of object returned by the server on registeration
type serverDelegate struct {
	// the number of workers assigned initially to the project
	workers int
	// Addressses of the workers
	workerAdresses []net.TCPAddr
	// id of the project this delegate represents
	projectID string
}

func (s serverDelegate) Workers() int {
	return s.workers
}

func (s serverDelegate) Message(worker int, message string, args interface{}, reply interface{}) error {
	// determine the ip address from the worker id and its port address
	var ipAddress, port string
	ipAddress = "192.168.1.100"
	port = "4600"
	client, err := rpc.DialHTTP("tcp", ipAddress+":"+port)
	if err != nil {
		return err
	}
	defer client.Close()
	return client.Call(message, args, reply)
}

// This is the type of objects
type ServerDelegator interface {
	/*
	   This call will be passed a single integer argument which is the count of workers available for the project.
	   The ids assigned to these workers will be 0 to worker-1 e.g. if workers == 10 then the ids' assigned to these 10 workers are 0,1,...9.
	   This call is made only once for the first time. Subsequent modifications to this list are communicated by WorkerDead and NewWorker.
	*/
	Workers() int

	/*
	   RPC communication is facilitated by this interface. It essentially represents the same interface provided by the
	   standard library for asynchronous RPC but provides a layer of abstraction hiding details about the worker and other
	   standard library objects by simply signalling wether the call succeeded or failed by a single buffered boolean channel.
	   It takes as argument the following :-
	       * worker - the worker to whom the communication is being initiated
	       * message - a string that acts as the message to send to the worker
	       * arg - A single argument to send to the worker
	       * reply - The reply from the worker will be stored here. It should always be a pointer. There is no checking but not doing
	           so will simply fail for obvious reasons.

	   It is a blocking call.
	*/
	Message(worker int, message string, arg interface{}, reply interface{}) chan bool
}

// This function registers the project notifier provided by the project and returns a ServerDelegator that the project can use
// to fetch important data and perform critical functions like messaging
func Register(p ProjectNotifier) ServerDelegator {

}
