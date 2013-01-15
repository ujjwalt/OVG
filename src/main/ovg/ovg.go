package ovg
import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"path"
	"labix.org/v2/mgo" 	
	"labix.org/v2/mgo/bson"
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

	NewWorker(newWorker int, project_id string, ip int)
}

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

	   A single result is returned - a singly buffered boolean channel which signals wether the request succeeded or not.
	*/

	Message(worker int, message string, arg interface{}, reply interface{}) chan bool
}

// The type of object returned by the server on registeration
type serverDelegate struct {
	
	workers int
	projectID string
}

type projectNotify struct {

	workers int
	projectID string
	ip int
}

func (s serverDelegate) Workers() int {
	var cnt int
	cnt := db.serverdb.distinct('workers')
	return cnt
}

func (s serverDelegate) Message(worker int, message string, arg interface{}, reply interface{}) chan bool {

	// Code to send message to host.
	//Why return workers ?
	return s.workers
}

// This function registers the project notifier provided by the project and returns a ServerDelegator that the project can use
// to fetch important data and perform critical functions like messaging


func Register(p ProjectNotify) ServerDelegator {
	db.serverdb.insert({p.workers, p.projectID, p.ip})		
}
	
func (p ProjectNotifier) WorkerDead(deadWorker int, replaced bool) {
	
} 
func (p ProjectNotifier) NewWorker(newWorker int, project_id string, ip int) {
	db.serverdb.insert({p.workers, p.projectID, p.ip})		
}
