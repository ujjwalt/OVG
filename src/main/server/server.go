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
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

//Worker counts the number of hosts
func workerCount() int
{
	return db.runCommand( { count: 'id' } )
}

// Returns an array of the distinct values of field id from all documents in the workers collection
func Worker(worker int) []string{
return db.runCommand ({ distinct: 'workers', key: 'id' } ) 
}

func Message(worker int, message string, args *Args , reply *int) chan bool {
	*reply="true"
if(*reply==nill){
	fmt.Println("Worker dead")
	//replaceWorker();
	}
return nil
}
func replaceWorker(worker_id int,project_id int)
{

}
		




