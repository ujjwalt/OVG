package ovg

type ProjectNotifier interface {
	
	WorkerDead(deadWorker int, replaced bool)
	NewWorker(newWorker int)
}

// The type of object returned by the server on registeration
type serverDelegate struct {
	
	workers int
	projectID string
}

func (s serverDelegate) Workers() int {
	s.workers= quer
	return s.workers
}

func (s serverDelegate) Message(worker int, message string, arg interface{}, reply interface{}) chan bool {
	s.workers= query
	return s.workers
}

type ServerDelegator interface {
	Workers() int
	Message(worker int, message string, arg interface{}, reply interface{}) chan bool
}


func Register(p ProjectNotifier) ServerDelegator {
	p.workers();
}
