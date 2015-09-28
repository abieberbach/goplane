package task
import "github.com/abieberbach/goplane/xplm/processing"

type task struct {
	taskFunction TaskFunc
	data         interface{}
	result       interface{}
	taskError    error
	doneChannel  chan bool
}

type TaskFunc func(data interface{}) (result interface{}, err error)

type TaskManager struct {
	taskChannel chan *task
}

func NewTaskManager(bufferSize int) *TaskManager {
	return &TaskManager{make(chan *task, bufferSize)}
}

func (self *TaskManager) Start() {
	processing.RegisterFlightLoopCallback(self.processTaskLoop, -1.0, nil)
}

func (self *TaskManager) Stop() {
	processing.UnregisterFlightLoopCallback(self.processTaskLoop, nil)
}
func (self *TaskManager) processTaskLoop(elapsedSinceLastCall, elapsedTimeSinceLastFlightLoop float32, counter int, ref interface{}) float32 {
	var req *task
	for {
		select {
		case req = <-self.taskChannel:
			req.result, req.taskError = req.taskFunction(req.data)
			req.doneChannel <- true
		default:
		//es gibt nichts zum verarbeiten
			return 0.0
		}
	}
	return 0.0
}

func (self *TaskManager) ExecuteTask(taskFunc TaskFunc, data interface{}) (result interface{}, err error) {
	doneChannel := make(chan bool)
	req := &task{taskFunc, data, nil, nil, doneChannel}
	self.taskChannel <- req
	processing.SetFlightLoopCallbackInterval(self.processTaskLoop, -1.0, true, nil)
	<-doneChannel
	return req.result, req.taskError
}