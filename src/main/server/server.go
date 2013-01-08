//server.go will provide the interface for communicating and handling hosts

// workerDead(message string), send the message and wait for ack, if no ack means worker dead

package 
main
import(
   
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
	"path"
	"os"
)

type Flag int
type Args struct{
	message string
}
func main() {

	flag := new(Flag)
	rpc.Register(flag)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":1234", nil) //ListenAndServe starts an HTTP server with a given address and handler. 
						//The handler is usually nil, which means to use DefaultServeMux. 
	if err != nil {
		fmt.Println(err.Error())
	}
}

//Worker counts the number of hosts
func workerCount() int
{
	return db.runCommand( { count: 'id' } ) //mongodb command to count the id
}

// Returns an array of the distinct values of field id from all documents in the workers collection
//In mongodb document is analogous to rdbms table and collection is record

func Worker(worker int) []string{
	return db.runCommand ({ distinct: 'workers', key: 'id' } ) //mongodb command to get the array of list of 
								   //workers for column id
}

func Message(worker int, message string, args *Args , reply *int) chan bool {
	
	
	Dial("tcp","192.168.23.12") //IP address of host 
	*reply="true"
if(*reply==nill){
	fmt.Println("Worker dead")
	//replaceWorker();
	}
return nil
}
//Replace dead worker of particular project with fresh worker
func replaceWorker(worker_id int,project_id int)
{
	db.workers.update(  //this query updates workers collection with id=worker_id(id of dead worker)
   { _id: worker_id, _project_id: project_id },
   {
     //$set: { ' ': 'Warner' },
   }
		
)
	
}
		




